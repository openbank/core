package core

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpctags "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tags"

	_ "github.com/lib/pq"

	"bnk.to/core/tools/auth"
	"bnk.to/core/tools/db/mux"
	"bnk.to/core/tools/services"
	"bnk.to/core/tools/services/backups"
	"bnk.to/core/tools/services/branches"
	"bnk.to/core/tools/services/bulks"
	"bnk.to/core/tools/services/cards"
	"bnk.to/core/tools/services/centres"
	"bnk.to/core/tools/services/clients"
	"bnk.to/core/tools/services/comments"
	"bnk.to/core/tools/services/consumers"
	"bnk.to/core/tools/services/currencies"
	"bnk.to/core/tools/services/deposits"
	"bnk.to/core/tools/services/documents"
	"bnk.to/core/tools/services/fields"
	"bnk.to/core/tools/services/groups"
	"bnk.to/core/tools/services/imports"
	"bnk.to/core/tools/services/jobs"
	"bnk.to/core/tools/services/ledgers"
	"bnk.to/core/tools/services/loans"
	"bnk.to/core/tools/services/metrics"
	"bnk.to/core/tools/services/notifications"
	"bnk.to/core/tools/services/org"
	"bnk.to/core/tools/services/products"
	"bnk.to/core/tools/services/rates"
	"bnk.to/core/tools/services/reports"
	"bnk.to/core/tools/services/revolving"
	"bnk.to/core/tools/services/roles"
	"bnk.to/core/tools/services/tasks"
	"bnk.to/core/tools/services/users"
	"bnk.to/core/tools/services/util"
)

// Server is the implementation of the core banking service.
type Server struct {
	logger zerolog.Logger
	addr   string
	auth   auth.Auth
	dbURL  string

	HealthCheck metrics.HealthCheckServer

	Backups         *backups.Server
	Branches        *branches.Server
	Bulks           *bulks.Server
	Cards           *cards.Server
	Centres         *centres.Server
	Clients         *clients.Server
	Comments        *comments.Server
	Consumers       *consumers.Server
	Currencies      *currencies.Server
	CurrencyRates   *currencies.RatesServer
	Deposits        *deposits.Server
	Documents       *documents.Server
	Fields          *fields.Server
	Groups          *groups.Server
	Imports         *imports.Server
	Jobs            *jobs.Server
	LedgerAccounts  *ledgers.AccountsServer
	LedgerEntries   *ledgers.EntriesServer
	Loans           *loans.Server
	Notifications   *notifications.Server
	Channels        *org.ChannelsServer
	Holidays        *org.HolidaysServer
	Org             *org.Server
	Accounting      *org.AccountingServer
	Setup           *org.SetupServer
	DepositProducts *products.DepositsServer
	LoanProducts    *products.LoansServer
	RiskLevels      *products.RiskLevelsServer
	Rates           *rates.Server
	Reports         *reports.Server
	Revolving       *revolving.AccountsServer
	ClientRoles     *roles.ClientsServer
	UserRoles       *roles.UsersServer
	Tasks           *tasks.Server
	Users           *users.Server
	Util            *util.Server
}

// NewServer initializes values in the core banking server.
func NewServer(logger zerolog.Logger, args Args) (*Server, error) {
	var authBackend auth.Auth
	switch args.AuthType {
	case "dummy":
		authBackend = auth.Dummy{}
	case "hydra":
		var err error
		authBackend, err = auth.NewHydra(args.HydraAdmin, logger)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid auth type: %s", args.AuthType)
	}
	common := services.Common{
		Auth: authBackend,
	}
	srv := &Server{
		logger: logger,
		addr:   args.Addr,
		auth:   authBackend,
		dbURL:  args.Database,

		Backups:         backups.NewServer(common),
		Branches:        branches.NewServer(common),
		Bulks:           bulks.NewServer(common),
		Cards:           cards.NewServer(common),
		Centres:         centres.NewServer(common),
		Clients:         clients.NewServer(common),
		Comments:        comments.NewServer(common),
		Consumers:       consumers.NewServer(common),
		Currencies:      currencies.NewServer(common),
		CurrencyRates:   currencies.NewRatesServer(common),
		Deposits:        deposits.NewServer(common),
		Documents:       documents.NewServer(common),
		Fields:          fields.NewServer(common),
		Groups:          groups.NewServer(common),
		Imports:         imports.NewServer(common),
		Jobs:            jobs.NewServer(common),
		LedgerAccounts:  ledgers.NewAccountsServer(common),
		LedgerEntries:   ledgers.NewEntriesServer(common),
		Loans:           loans.NewServer(common),
		Notifications:   notifications.NewServer(common),
		Channels:        org.NewChannelsServer(common),
		Holidays:        org.NewHolidaysServer(common),
		Org:             org.NewServer(common),
		Accounting:      org.NewAccountingServer(common),
		Setup:           org.NewSetupServer(common),
		DepositProducts: products.NewDepositsServer(common),
		LoanProducts:    products.NewLoansServer(common),
		RiskLevels:      products.NewRiskLevelsServer(common),
		Rates:           rates.NewServer(common),
		Reports:         reports.NewServer(common),
		Revolving:       revolving.NewAccountsServer(common),
		ClientRoles:     roles.NewClientsServer(common),
		UserRoles:       roles.NewUsersServer(common),
		Tasks:           tasks.NewServer(common),
		Users:           users.NewServer(common),
		Util:            util.NewServer(common),
	}
	return srv, nil
}

// Run starts the core banking server, registering all services.
func (s *Server) Run(ctx context.Context) error {
	// listen
	s.logger.Info().Str("addr", s.addr).Msg("listen")
	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", s.addr)
	if err != nil {
		return fmt.Errorf("error listening on %q: %w", s.addr, err)
	}
	defer l.Close()
	// serve
	return s.grpcServer().Serve(l)
}

// grpcServer creates the grpc server.
func (s *Server) grpcServer() *grpc.Server {
	// build server
	srv := grpc.NewServer(
		grpcmw.WithUnaryServerChain(
			grpctags.UnaryServerInterceptor(grpctags.WithFieldExtractor(grpctags.CodeGenRequestFieldExtractor)),
			grpclog.UnaryServerInterceptor(grpczerolog.InterceptorLogger(s.logger)),
			mux.Interceptor(s.dbURL, s.auth),
		),
		grpcmw.WithStreamServerChain(
			grpctags.StreamServerInterceptor(grpctags.WithFieldExtractor(grpctags.CodeGenRequestFieldExtractor)),
			grpclog.StreamServerInterceptor(grpczerolog.InterceptorLogger(s.logger)),
		),
	)
	// add reflection, health, util
	reflection.Register(srv)
	s.registerServices(srv)
	return srv
}
