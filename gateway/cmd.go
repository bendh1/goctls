package gateway

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/duke-git/lancet/v2/fileutil"

	"github.com/bendh1/goctls/internal/cobrax"
	"github.com/bendh1/goctls/util/ctx"
	"github.com/bendh1/goctls/util/pathx"
	"github.com/spf13/cobra"
)

var (
	varStringHome   string
	varStringRemote string
	varStringBranch string
	varStringDir    string
	varStringName   string
	varIntPort      int

	Cmd = cobrax.NewCommand("gateway", cobrax.WithRunE(generateGateway))
)

func init() {
	CmdFlags := Cmd.Flags()

	CmdFlags.StringVar(&varStringHome, "home")
	CmdFlags.StringVar(&varStringRemote, "remote")
	CmdFlags.StringVar(&varStringBranch, "branch")
	CmdFlags.StringVarPWithDefaultValue(&varStringDir, "dir", "d", ".")
	CmdFlags.StringVarPWithDefaultValue(&varStringName, "name", "n", "gateway")
	CmdFlags.IntVarPWithDefaultValue(&varIntPort, "port", "p", 8080)
}

func generateGateway(*cobra.Command, []string) error {
	path, err := filepath.Abs(varStringDir)
	if err != nil {
		return err
	}

	path = filepath.Join(path, varStringName)

	err = fileutil.CreateDir(path)
	if err != nil {
		return err
	}

	if _, err := ctx.Prepare(path); err != nil {
		return err
	}

	etcContent, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	mainContent, err := pathx.LoadTemplate(category, mainTemplateFile, mainTemplate)
	if err != nil {
		return err
	}

	etcDir := filepath.Join(path, "etc")
	if err := pathx.MkdirIfNotExist(etcDir); err != nil {
		return err
	}
	etcFile := filepath.Join(etcDir, "gateway.yaml")

	etcTpl, err := template.New("etc").Parse(etcContent)
	if err != nil {
		return err
	}

	var etcData strings.Builder
	err = etcTpl.Execute(&etcData, map[string]any{
		"port": varIntPort,
		"name": varStringName,
	})

	err = fileutil.WriteStringToFile(etcFile, etcData.String(), false)
	if err != nil {
		return err
	}

	mainFile := filepath.Join(path, "main.go")
	return os.WriteFile(mainFile, []byte(mainContent), 0644)
}
