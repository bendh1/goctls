package gogen

import (
	_ "embed"
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bendh1/goctls/api/spec"
	"github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/util/format"
	"github.com/bendh1/goctls/util/pathx"
	"github.com/bendh1/goctls/vars"
)

const contextFilename = "service_context"

//go:embed svc.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, api *spec.ApiSpec, g *GenContext) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}

	var middlewareStr string
	var middlewareAssignment string
	middlewares := getMiddleware(api)

	for _, item := range middlewares {
		middlewareStr += fmt.Sprintf("%s rest.Middleware\n", item)
		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		middlewareAssignment += fmt.Sprintf("%s: %s,\n", item,
			fmt.Sprintf("middleware.New%s().%s", cases.Title(language.English, cases.NoLower).String(name), "Handle"))
	}

	configImport := "\"" + pathx.JoinPackages(rootPkg, configDir) + "\""
	if len(middlewareStr) > 0 {
		configImport += "\n\t\"" + pathx.JoinPackages(rootPkg, middlewareDir) + "\""
		configImport += fmt.Sprintf("\n\t\"%s/rest\"", vars.ProjectOpenSourceURL)
	}

	if g.UseI18n {
		configImport += fmt.Sprintf("\n\ti18n2 \"%s/internal/i18n\"", rootPkg)
	}

	if g.UseCasbin {
		configImport += fmt.Sprintf("\n\t\"%s/internal/middleware\"", rootPkg)
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          contextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]any{
			"configImport":         configImport,
			"config":               "config.Config",
			"middleware":           middlewareStr,
			"hasMiddleware":        len(middlewares) > 3,
			"middlewareAssignment": middlewareAssignment,
			"useCasbin":            g.UseCasbin,
			"useI18n":              g.UseI18n,
			"useEnt":               g.UseEnt,
			"projectPackage":       rootPkg,
			"useTrans":             g.TransErr,
			"useCoreRpc":           g.UseCoreRpc,
		},
	})
}
