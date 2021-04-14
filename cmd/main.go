package main

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/commands"
	"github.com/spf13/cobra"
	"log"
)

var (
	version = "v0.0.1"

	rootCmd = &cobra.Command{
		Use:     "conflict",
		Short:   "conflict",
		Long:    `conflict`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(commands.NewNewSiteCmd, commands.NewBuildCmd, commands.NewServeCmd, commands.NewNewFile, commands.NewExec)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	//service.Init()
	//service.Run()
}
