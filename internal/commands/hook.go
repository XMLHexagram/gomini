package commands

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "",
		Long:  "",
		RunE:  hook,
	}

	//cmd.Flags().Bool("force", false, "force or not")
	NewHook = cmd
}

var NewHook = &cobra.Command{}

func hook(cmd *cobra.Command, args []string) error {
	ginS.POST("/gomini/hook/", func(c *gin.Context) {
		exec.Command("git","pull")
	})
	err := ginS.Run(":12210")
	if err != nil {
		panic(err)
	}
	return nil
}
