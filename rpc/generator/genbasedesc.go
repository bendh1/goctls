package generator

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	conf "github.com/bendh1/goctls/config"
	"github.com/bendh1/goctls/rpc/parser"
	"github.com/bendh1/goctls/util"
	"github.com/bendh1/goctls/util/pathx"
)

func (g *Generator) GenBaseDesc(ctx DirContext, _ parser.Proto, cfg *conf.Config, c *ZRpcContext) error {
	protoFilename := filepath.Join(ctx.GetMain().Filename, "desc", "base.proto")
	if err := pathx.MkdirIfNotExist(filepath.Join(ctx.GetMain().Filename, "desc")); err != nil {
		return err
	}

	err := util.With("t").Parse(rpcTemplateText).SaveTo(map[string]string{
		"package":     strings.ToLower(c.RpcName),
		"serviceName": strcase.ToCamel(c.RpcName),
	}, protoFilename, false)
	return err
}
