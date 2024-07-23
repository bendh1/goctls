package schema

import (
	_ "embed"
	"errors"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bendh1/goctls/util/format"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	VarStringModelName string
)

//go:embed tpl/basic.tpl
var schemaTpl string

func GenSchema(_ *cobra.Command, _ []string) error {
	if VarStringModelName == "" {
		return errors.New("the model name can not be empty")
	}

	var schemaStr strings.Builder

	schemaTpl, err := template.New("schemaTpl").Parse(schemaTpl)
	if err != nil {
		return err
	}

	err = schemaTpl.Execute(&schemaStr, map[string]string{
		"ModelName":          VarStringModelName,
		"ModelNameLowercase": strings.ToLower(VarStringModelName),
	})

	if err != nil {
		return err
	}

	var filePath string
	tmp, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	if strings.HasSuffix(tmp, "schema") {
		filePath = tmp
	} else {
		newPath := filepath.Join(tmp, "ent/schema")
		if fileutil.IsExist(newPath) {
			filePath = newPath
		} else {
			entpath := filepath.Join(tmp, "ent")
			if fileutil.IsExist(entpath) {
				err = fileutil.CreateDir(newPath)
				if err != nil {
					return err
				}
			} else {
				return errors.New("failed to find the ent schema folder")
			}
		}
	}

	filename, err := format.FileNamingFormat("go_zero", VarStringModelName)
	if err != nil {
		return err
	}

	err = fileutil.WriteStringToFile(filepath.Join(filePath, filename+".go"),
		schemaStr.String(), false)
	if err != nil {
		return err
	}

	color.Green.Println("Generate Ent schema successfully")

	return nil
}
