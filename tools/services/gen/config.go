package main

// modelOverride overrides ModelName, mapping from the method name without
// the verb to the expected model name.
var modelOverride = map[string][2]string{
	"deposits.Card": {"Card", "Cards"},
	"loans.Card":    {"Card", "Cards"},
	"loans.Funding": {"LoanInvestorFund", "LoanInvestorFunds"},
}

// serverOverride overrides ServerName, mapping from the name of the
// service to the expected server name.
var serverOverride = map[string]string{
	"backups.DatabaseBackupService":   "Server",
	"org.OrganizationService":         "Server",
	"roles.ClientRoleService":         "ClientsServer",
	"roles.UserRolesService":          "UsersServer",
	"products.DepositProductsService": "DepositsServer",
	"products.LoanProductsService":    "LoansServer",
}

// uniqueOverride overrides the name of the ID field of a given struct.
var uniqueOverride = map[string]string{
	"CardHoldConfig": "MerchantCode",
	"Currency":       "Code",
}
