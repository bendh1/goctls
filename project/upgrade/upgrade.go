package upgrade

import (
	"errors"
	"os"

	conf "github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/rpc/execx"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	// VarBoolUpgradeMakefile describe whether to upgrade makefile
	VarBoolUpgradeMakefile bool
)

func UpgradeProject(_ *cobra.Command, _ []string) error {
	color.Green.Println("Start upgrading dependencies...")

	err := editMod(conf.DefaultGoZeroVersion, conf.DefaultToolVersion)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = upgradeDependencies(wd)
	if err != nil {
		return err
	}

	if VarBoolUpgradeMakefile {
		color.Green.Println("Start upgrading Makefile ...")
		_, err = execx.Run("goctls extra makefile", wd)
		if err != nil {
			return errors.New("failed to upgrade makefile")
		}
	}

	color.Green.Println("Done.")

	return nil
}
