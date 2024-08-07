package gogen

import (
	_ "embed"
	"path/filepath"

	"github.com/bendh1/goctls/util/format"
	"github.com/bendh1/goctls/util/pathx"
)

//go:embed enttx.tpl
var entTxTemplateText string

func GenEntTx(dir, style, rootPkg string) error {
	filename, err := format.FileNamingFormat(style, "ent_tx.go")
	if err != nil {
		return err
	}

	if err := pathx.MkdirIfNotExist(filepath.Join(dir, utilsDir, "entx")); err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          filepath.Join(utilsDir, "entx"),
		filename:        filename,
		templateName:    "entTxFile",
		category:        category,
		templateFile:    entTxTemplateFile,
		builtinTemplate: entTxTemplateText,
		data: map[string]string{
			"package": rootPkg,
		},
	})
}
