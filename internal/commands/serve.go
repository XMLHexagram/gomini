package commands

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-server/pkg/gemini"
	"github.com/spf13/cobra"
)

var NewServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long:  "",
	RunE:  newServe,
}

func newServe(cmd *cobra.Command, args []string) error {
	config.Init()
	tls := config.ProvideTLS()
	base := config.ProvideBase()
	engine, err := gemini.New(tls.CertFile, tls.KeyFile, base.LanguageCode)
	if err != nil {
		panic(err)
	}
	engine.HandleDir("/", "public", base.Entry)
	err = engine.Run(":1965")
	if err != nil {
		panic(err)
	}
	return nil
}
