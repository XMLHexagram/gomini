package commands

import (
	"fmt"
	"github.com/lmx-Hexagram/gemini-generator/pkg"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var _ cmder = (*newSiteCmd)(nil)

type newSiteCmd struct {
	//configFormat string

	*baseBuilderCmd
}

func (b *commandsBuilder) newNewSiteCmd() *newSiteCmd {
	cc := &newSiteCmd{}

	cmd := &cobra.Command{
		Use:   "site [path]",
		Short: "Create a new site (skeleton)",
		Long: `Create a new site in the provided directory.
The new site will have the correct structure, but no content or theme yet.
Use ` + "`gomini new [contentPath]`" + ` to create new content.`,
		RunE: cc.newSite,
	}

	//cmd.Flags().StringVarP(&cc.configFormat, "format", "f", "toml", "config & frontmatter format")
	//cmd.Flags().Bool("force", false, "init inside non-empty directory")

	cc.baseBuilderCmd = b.newBuilderCmd(cmd) // todo: basic

	return cc
}

//var NewNewSiteCmd = &cobra.Command{
//	Use:   "new-site",
//	Short: "Create a gomini template",
//	Long:  "",
//	RunE:  newSite,
//}

func (n *newSiteCmd) newSite(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: project name is required.\033[m Example: gomini new site helloworld\n")
		return nil
	}
	fmt.Println()
	createpath, err := filepath.Abs(filepath.Clean(args[0]))
	//pkg.Replaces[0].After = args[0]
	if err != nil {
		return err
	}
	err = pkg.GenerateDir("template", createpath)
	if err != nil {
		return err
	}

	return nil
}
