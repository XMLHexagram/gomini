package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ cmder = (*versionCmd)(nil)

type versionCmd struct {
	*baseCmd
}

func newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("version 0.0.3")
			return nil
		},
	}

	cc.baseCmd = newBaseCmd(cmd)
	return cc
}
