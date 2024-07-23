package gogen

import (
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/rpc/execx"
	"github.com/duke-git/lancet/v2/fileutil"
)

func genCasbin(dir string, cfg *config.Config, g *GenContext) error {
	var useI18n string
	if g.TransErr {
		useI18n = " -i"
	}

	if !fileutil.IsExist(filepath.Join(dir, middlewareDir)) {
		err := fileutil.CreateDir(filepath.Join(dir, middlewareDir))
		if err != nil {
			return err
		}
	}

	_, err := execx.Run(fmt.Sprintf("goctls extra middleware api -a authority%s -s %s", useI18n, cfg.NamingFormat), dir)

	if err != nil {
		return err
	}

	return nil
}
