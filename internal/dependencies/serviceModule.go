package dependencies

import "github.com/xandervanderweken/Password/internal/services"

var GeneratorService = services.NewGeneatorService()
var AccountEntryService = services.NewAccountService(AccountEntryRepository)
