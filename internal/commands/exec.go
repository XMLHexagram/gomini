package commands

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

//openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out MyCertificate.crt -keyout MyKey.key

func init() {
	cmd := &cobra.Command{
		Use:   "exec",
		Short: "exec scrpit",
		Long:  "",
		RunE:  newExec,
	}

	//cmd.Flags().String("path", "", "where in script")
	NewExec = cmd
}

var NewExec = &cobra.Command{}

func newExec(cmd *cobra.Command, args []string) error {
	//path, err := cmd.Flags().GetString("path")
	//if err != nil {
	//	return err
	//}

	truePath := filepath.Join("script", args[0])
	c := exec.Command("/bin/bash", truePath)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
