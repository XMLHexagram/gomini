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
	geminiConfig := config.ProvideGemini()
	engine, err := gemini.New(tls.CertFile, tls.KeyFile, geminiConfig.DefaultLang)
	if err != nil {
		panic(err)
	}
	engine.AutoRedirect = geminiConfig.AutoRedirect
	engine.AutoRedirectUrl = geminiConfig.AutoRedirectUrl
	for _, v := range geminiConfig.Dir {
		engine.HandleDir(v.Router, v.Path, v.Index)
	}

	for _, v := range geminiConfig.File {
		engine.HandleFile(v.Router, v.Path)
	}

	for _, v := range geminiConfig.Proxy {
		engine.HandleProxy(v.Router, v.URL)
	}
	err = engine.Run(":1965")
	if err != nil {
		panic(err)
	}
	return nil
}
