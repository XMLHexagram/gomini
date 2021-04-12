package commands

import (
	"fmt"
	"github.com/lmx-Hexagram/gemini-generator/pkg"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var NewNewSiteCmd = &cobra.Command{
	Use:   "new-site",
	Short: "Create a gomini template",
	Long:  "",
	RunE:  newSite,
}

func newSite(cmd *cobra.Command, args []string) error {
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
