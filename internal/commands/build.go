package commands

import (
	"fmt"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/spf13/cobra"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var _ cmder = (*newBuild)(nil)

type newBuild struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newNewBuild() *newBuild {
	cc := &newBuild{}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "Create new content for your site",
		Long: `Create a new content file and automatically set the date and title.
It will guess which kind of file to create based on the path provided.

You can also specify the kind with ` + "`-k KIND`" + `.

If archetypes are provided in your theme or site, they will be used.

Ensure you run this within the root directory of your site.`,
		RunE: build,
	}

	cc.baseBuilderCmd = b.newBuilderCmd(cmd)

	return cc
}

func build(cmd *cobra.Command, args []string) error {
	config.Init()
	baseConfig := config.ProvideBase()
	fmt.Println(baseConfig)
	type templateData struct {
		BaseUrl string
	}
	var td = templateData{BaseUrl: baseConfig.BaseUrl}
	createpath, err := filepath.Abs(filepath.Clean("public"))
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(createpath, 0777)
	if err != nil {
		panic(err)
	}
	filepath.WalkDir("./", func(path_ string, d fs.DirEntry, err error) error {
		//a, b := filepath.Split(path_)
		//fmt.Println(path_)
		//fmt.Println(a, b)
		//filter := func(a string) bool{
		//fmt.Println(path_,":::::")
		if ok, err := path.Match("content/*", path_); !ok || err != nil {
			//fmt.Println(path_,"::")
			return nil
		}
		//}
		//var blocks = []string{".git",""}

		pathList := strings.Split(path_, "/")
		//fmt.Println(pathList)
		if pathList[0] == "public" || path_ == "./" || len(pathList) == 0 || string(pathList[len(pathList)-1][0]) == "." {
			return nil
		}
		if d.IsDir() == true {
			pathList[0] = "public"
			os.MkdirAll(path.Join(pathList...), 0777)
		} else {
			if len(pathList) == 1 {
				fmt.Println(pathList)
				return nil
			}
			var temp = pathList
			temp[0] = "public"
			f, err := os.Create(path.Join(temp...))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			//fmt.Println(path_)
			tmpl, err := template.ParseFiles(path_)
			if err != nil {
				panic(err)
			}
			//tmpl.Execute(os.Stdout, td)
			tmpl.Execute(io.Writer(f), td)
		}
		return nil
	})
	//os.ReadDir()
	//dirFs := os.DirFS("")

	//tmpl, err := template.ParseFS(dirFs)
	//if err != nil {
	//	return err
	//}
	//tmpl.ParseFiles()
	return nil
}

var ignore = []string{".DS_Store", "template.go", "template_test.go", ".git"}
var normalFileMode = os.FileMode(0644)
var normalDirMode = os.FileMode(0755)

type Replace struct {
	Before string
	After  string
}

var Replaces = []Replace{
	{

	},
}
