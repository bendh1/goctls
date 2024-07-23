package upgrade

import (
	"fmt"

	"github.com/bendh1/goctls/rpc/execx"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// upgrade gets the latest goctl by
// go install github.com/bendh1/goctls@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `go install github.com/bendh1/goctls@latest`
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	color.Green.Println("Upgrade successfully")
	return nil
}
