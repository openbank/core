package core

import (
	"bnk.to/core/api/v1/backups"
	"bnk.to/core/api/v1/branches"
	"bnk.to/core/api/v1/bulks"
	"bnk.to/core/api/v1/cards"
	"bnk.to/core/api/v1/centres"
	"bnk.to/core/api/v1/clients"
	"bnk.to/core/api/v1/comments"
	"bnk.to/core/api/v1/consumers"
	"bnk.to/core/api/v1/currencies"
	"bnk.to/core/api/v1/deposits"
	"bnk.to/core/api/v1/documents"
	"bnk.to/core/api/v1/fields"
	"bnk.to/core/api/v1/groups"
	"bnk.to/core/api/v1/imports"
	"bnk.to/core/api/v1/jobs"
	"bnk.to/core/api/v1/ledgers"
	"bnk.to/core/api/v1/notifications"
	"bnk.to/core/api/v1/org"
	"bnk.to/core/api/v1/products"
	"bnk.to/core/api/v1/rates"
	"bnk.to/core/api/v1/reports"
	"bnk.to/core/api/v1/revolving"
	"bnk.to/core/api/v1/roles"
	"bnk.to/core/api/v1/tasks"
	"bnk.to/core/api/v1/users"
	"bnk.to/core/api/v1/util"
	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health/grpc_health_v1"
)

func (s *Server) registerServices(grpcServer *grpc.Server) {
	grpchealth.RegisterHealthServer(grpcServer, s.HealthCheck)

	backups.RegisterBackupServiceServer(grpcServer, s.Backups)
	branches.RegisterBranchesServiceServer(grpcServer, s.Branches)
	bulks.RegisterBulksServiceServer(grpcServer, s.Bulks)
	cards.RegisterCardsServiceServer(grpcServer, s.Cards)
	centres.RegisterCentresServiceServer(grpcServer, s.Centres)
	clients.RegisterClientsServiceServer(grpcServer, s.Clients)
	comments.RegisterCommentsServiceServer(grpcServer, s.Comments)
	consumers.RegisterConsumersServiceServer(grpcServer, s.Consumers)
	currencies.RegisterCurrencyServiceServer(grpcServer, s.Currencies)
	currencies.RegisterRatesServiceServer(grpcServer, s.CurrencyRates)
	deposits.RegisterDepositsServiceServer(grpcServer, s.Deposits)
	documents.RegisterDocumentsServiceServer(grpcServer, s.Documents)
	fields.RegisterFieldsServiceServer(grpcServer, s.Fields)
	groups.RegisterGroupsServiceServer(grpcServer, s.Groups)
	imports.RegisterImportsServiceServer(grpcServer, s.Imports)
	jobs.RegisterJobsServiceServer(grpcServer, s.Jobs)
	ledgers.RegisterAccountsServiceServer(grpcServer, s.LedgerAccounts)
	ledgers.RegisterEntriesServiceServer(grpcServer, s.LedgerEntries)
	notifications.RegisterNotificationsServiceServer(grpcServer, s.Notifications)
	org.RegisterChannelsServiceServer(grpcServer, s.Channels)
	org.RegisterHolidaysServiceServer(grpcServer, s.Holidays)
	org.RegisterOrganizationServiceServer(grpcServer, s.Org)
	org.RegisterAccountingServiceServer(grpcServer, s.Accounting)
	org.RegisterSetupServiceServer(grpcServer, s.Setup)
	products.RegisterDepositProductsServiceServer(grpcServer, s.DepositProducts)
	products.RegisterLoanProductsServiceServer(grpcServer, s.LoanProducts)
	products.RegisterRiskLevelsServiceServer(grpcServer, s.RiskLevels)
	rates.RegisterRatesServiceServer(grpcServer, s.Rates)
	reports.RegisterReportsServiceServer(grpcServer, s.Reports)
	revolving.RegisterAccountsServiceServer(grpcServer, s.Revolving)
	roles.RegisterClientRoleServiceServer(grpcServer, s.ClientRoles)
	roles.RegisterUserRolesServiceServer(grpcServer, s.UserRoles)
	tasks.RegisterTasksServiceServer(grpcServer, s.Tasks)
	users.RegisterUsersServiceServer(grpcServer, s.Users)
	util.RegisterUtilServiceServer(grpcServer, s.Util)
}
