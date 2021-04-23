package commands

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/create"
	"github.com/spf13/cobra"
)

var _ cmder = (*newCmd)(nil)

type newCmd struct {
	contentEditor string
	contentType   string

	*baseBuilderCmd
}

func (b *commandsBuilder) newNewCmd() *newCmd {
	cmd := &cobra.Command{
		Use:   "new [path]",
		Short: "Create new content for your site",
		Long: `
Create a new content.
Ensure you run this within the root directory of your site.`,
		RunE: newFile,
	}

	cc := &newCmd{baseBuilderCmd: b.newBuilderCmd(cmd)}

	cmd.AddCommand(b.newNewSiteCmd().getCommand())
	return cc
}

func newFile(cmd *cobra.Command, args []string) error {
	//isForce, err := cmd.Flags().GetBool("force")
	//if err != nil {
	//	return err
	//}
	createPath := args[0]

	return create.NewContent(createPath)
}
