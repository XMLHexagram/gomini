package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (b *commandsBuilder) newGominiCmd() *gominiCmd {
	cc := &gominiCmd{}

	cmd := &cobra.Command{
		Use:   "gomini",
		Short: "gomini build your Gemini site",
		Long: ` 
gomini is a cmd-cli, used to build your Gemini site.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Gemini // Gomini")
			return nil
		},
	}

	cc.baseBuilderCmd = b.newBuilderCmd(cmd)

	return cc
}
