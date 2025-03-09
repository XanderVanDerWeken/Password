package dependencies

import "github.com/xandervanderweken/Password/internal/controllers"

var PasswordController = controllers.NewPasswordController(GeneratorService)
