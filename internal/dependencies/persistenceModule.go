package dependencies

import "github.com/xandervanderweken/Password/internal/persistence"

var Db = persistence.LoadDatabase(ConfigObj.ConnectionString)
var AccountEntryRepository = persistence.NewAccountRepository(Db)
