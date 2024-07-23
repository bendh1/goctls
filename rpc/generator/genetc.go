package generator

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	conf "github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/rpc/parser"
	"github.com/bendh1/goctls/util"
	"github.com/bendh1/goctls/util/format"
	"github.com/bendh1/goctls/util/pathx"
	"github.com/bendh1/goctls/util/stringx"
)

//go:embed etc.tpl
var etcTemplate string

// GenEtc generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEtc(ctx DirContext, _ parser.Proto, cfg *conf.Config, c *ZRpcContext) error {
	dir := ctx.GetEtc()
	etcFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, fmt.Sprintf("%v.yaml", etcFilename))

	text, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	return util.With("etc").Parse(text).SaveTo(map[string]any{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
		"isEnt":       c.Ent,
		"port":        c.Port,
	}, fileName, false)
}
