package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/bendh1/goctls/cicd"
	"github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/project"
	"github.com/bendh1/goctls/upgrade"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	cobracompletefig "github.com/withfig/autocomplete-tools/integrations/cobra"

	"github.com/bendh1/goctls/api"
	"github.com/bendh1/goctls/docker"
	"github.com/bendh1/goctls/env"
	"github.com/bendh1/goctls/extra"
	"github.com/bendh1/goctls/frontend"
	"github.com/bendh1/goctls/gateway"
	"github.com/bendh1/goctls/info"
	"github.com/bendh1/goctls/internal/cobrax"
	"github.com/bendh1/goctls/internal/version"
	"github.com/bendh1/goctls/kube"
	"github.com/bendh1/goctls/rpc"
	"github.com/bendh1/goctls/tpl"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var (
	//go:embed usage.tpl
	usageTpl string
	rootCmd  = cobrax.NewCommand("goctls")
)

// Execute executes the given command
func Execute() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(codeFailure)
	}
}

func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}

func init() {
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    blue,
		"green":   green,
		"rpadx":   rpadx,
		"rainbow": rainbow,
	})

	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s - Go Zero %s - Simple Admin Tools %s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH, config.DefaultGoZeroVersion, config.DefaultToolVersion)

	rootCmd.SetUsageTemplate(usageTpl)
	rootCmd.AddCommand(api.Cmd, docker.Cmd, kube.Cmd, env.Cmd, gateway.Cmd)
	rootCmd.AddCommand(rpc.Cmd, tpl.Cmd, frontend.Cmd, extra.ExtraCmd, info.Cmd, upgrade.Cmd, cicd.CicdCmd,
		project.Cmd)
	rootCmd.Command.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
	rootCmd.MustInit()
}
