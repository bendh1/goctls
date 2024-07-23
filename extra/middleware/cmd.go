package middleware

import (
	"github.com/bendh1/goctls/extra/middleware/api"
	"github.com/bendh1/goctls/internal/cobrax"
)

var (
	MiddlewareCmd = cobrax.NewCommand("middleware")
	ApiCmd        = cobrax.NewCommand("api", cobrax.WithRunE(api.Gen))
)

func init() {
	ApiCmdFlags := ApiCmd.Flags()

	ApiCmdFlags.StringVarP(&api.VarStringName, "name", "a")
	ApiCmdFlags.StringVarPWithDefaultValue(&api.VarStringOutput, "output", "o", ".")
	ApiCmdFlags.BoolVarP(&api.VarBoolList, "list", "l")
	ApiCmdFlags.BoolVarP(&api.VarBoolI18n, "i18n", "i")
	ApiCmdFlags.StringVarPWithDefaultValue(&api.VarStringStyle, "style", "s", "go_zero")

	MiddlewareCmd.AddCommand(ApiCmd)
}
