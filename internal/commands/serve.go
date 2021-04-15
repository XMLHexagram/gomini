package commands

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-generator/pkg/gemini/gemini"
	"github.com/spf13/cobra"
	"os/exec"
)

type ServeWorker struct {
}

var NewServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long:  "",
	RunE:  newServe,
}

func newServe(cmd *cobra.Command, args []string) error {
	initProxy()

	channal := make(chan bool, 0)
	go listenHook(channal)
	for {
		err := build(nil, []string{})
		if err != nil {
			fmt.Println(err)
		}
		close := test()
		close2 := proxy()
		//go hook()

		<-channal

		close2()
		close()
		output, err := exec.Command("/bin/sh", "git pull").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(output))
	}

	return nil
}

func test() func() error {
	config.Init()
	tls := config.ProvideTLS()

	//context.WithCancel()
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
	close, err := engine.Run(":1965")
	if err != nil {
		panic(err)
	}
	return close
}

func listenHook(channal chan bool) {
	ginS.POST("/gomini/hook/", func(c *gin.Context) {
		channal <- true
	})
	err := ginS.Run(":12210")
	if err != nil {
		fmt.Println(err)
	}
}
