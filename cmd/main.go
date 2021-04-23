package main

import (
	"fmt"
	"github.com/lmx-Hexagram/gemini-generator/internal/commands"
	"os"
)

//var (
//	version = "v0.0.1"
//
//	rootCmd = &cobra.Command{
//		Use:     "gomini",
//		Short:   "",
//		Long:    ``,
//		Version: version,
//	}
//)

//func init() {
//	rootCmd.AddCommand(commands.NewNewSiteCmd, commands.NewBuildCmd, commands.NewServeCmd, commands.NewNewFile, commands.NewExec, commands.NewHook)
//}

func main() {
	resp := commands.Execute(os.Args[1:])
	if resp.Err != nil {
		fmt.Printf("%+v\n", resp)
		os.Exit(-1)
	}
	//service.Init()
	//service.Run()
}
