package commands

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-generator/pkg/gemini/gemini"
	"github.com/spf13/cobra"
)

type serverCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newServerCmd() *serverCmd {
	return b.newServerCmdSignaled(nil)
}

func (b *commandsBuilder) newServerCmdSignaled(stop <-chan bool) *serverCmd {
	cc := &serverCmd{}

	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"serve"},
		Short:   "A high performance webserver",
		Long:    `Gomini`,
		RunE:    cc.newServe,
	}
	cc.baseBuilderCmd = b.newBuilderCmd(cmd)
	return cc
}

func (cc *serverCmd) newServe(cmd *cobra.Command, args []string) error {
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
		r, err := git.PlainOpen(".")
		if err != nil {
			fmt.Println(err)
		}
		worktree, err := r.Worktree()
		if err != nil {
			fmt.Println(err)
		}
		err = worktree.Reset(&git.ResetOptions{
			Commit: plumbing.Hash{},
			Mode:   git.HardReset,
		})
		if err != nil {
			fmt.Println("newServe", err)
		}
		err = worktree.Clean(&git.CleanOptions{
			Dir: false,
		})
		if err != nil {
			fmt.Println("newServe", err)
		}
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			fmt.Println("newServe", err)
		}
		fmt.Println("pull success")
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
