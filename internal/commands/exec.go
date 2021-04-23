package commands

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var _ cmder = (*newExec)(nil)

type newExec struct {
	*baseBuilderCmd
}

//openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out MyCertificate.crt -keyout MyKey.key

func (b *commandsBuilder) newNewExec() *newExec {
	cc := &newExec{}

	cmd := &cobra.Command{
		Use:   "exec",
		Short: "exec",
		Long:  ``,
		RunE:  cc.Exec,
	}

	cc.baseBuilderCmd = b.newBuilderCmd(cmd)

	return cc
}

func (cc *newExec) Exec(cmd *cobra.Command, args []string) error {
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
