package commands

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a gomini file",
		Long:  "",
		RunE:  newFile,
	}

	cmd.Flags().Bool("force", false, "force or not")
	NewNewFile = cmd
}

var NewNewFile = &cobra.Command{}

func newFile(cmd *cobra.Command, args []string) error {
	isForce, err := cmd.Flags().GetBool("force")
	if err != nil {
		return err
	}
	userPath := args[0]
	path_ := filepath.Join("content", userPath)
	if isForce {
		dirPath := filepath.Dir(path_)
		err := os.MkdirAll(dirPath, 0766)
		if err != nil {
			return err
		}
	}
	if filepath.Ext(path_) == "" {
		path_ += ".gmi"
	}
	f, err := os.Create(path_)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
