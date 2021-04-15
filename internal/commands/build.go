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

var NewBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Create a gomini template",
	Long:  "",
	RunE:  build,
}

func build(cmd *cobra.Command, args []string) error {
	//if len(args) == 0 {
	//	fmt.Fprintf(os.Stderr, "\033[31mERROR: project name is required.\033[m Example: gomini new site helloworld\n")
	//	return nil
	//}
	//createpath, err := filepath.Abs(filepath.Clean(args[0]))
	//pkg.Replaces[0].After = args[0]
	//if err != nil {
	//	return err
	//}
	//err = pkg.GenerateDir("template", createpath)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
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
		pathList := strings.Split(path_, "/")
		fmt.Println(pathList)
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

			fmt.Println(path_)
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

var ignore = []string{".DS_Store", "template.go", "template_test.go"}
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

//func GenerateDir(template fs.FS,src, dst string) error {
//	if err := os.MkdirAll(dst, normalDirMode); err != nil {
//		return err
//	}
//
//	dirs, err := template.Open()
//		ReadDir(src)
//	if err != nil {
//		return err
//	}
//
//	for _, v := range dirs {
//		if hasSets(v.Name(), ignore) {
//			continue
//		}
//
//		srcfp := path.Join(src, v.Name())
//		dstfp := path.Join(dst, v.Name())
//
//		if v.IsDir() {
//			err = GenerateDir(srcfp, dstfp)
//			if err != nil {
//				return err
//			}
//		} else {
//			err := GenerateFile(srcfp, dstfp)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func GenerateFile(src, dst string) error {
//	buf, err := Template.ReadFile(src)
//	if err != nil {
//		return err
//	}
//	for _, v := range Replaces {
//		buf = bytes.ReplaceAll(buf, []byte(v.Before), []byte(v.After))
//	}
//	err = ioutil.WriteFile(dst, buf, normalFileMode)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func hasSets(name string, sets []string) bool {
//	for _, ig := range sets {
//		if ig == name {
//			return true
//		}
//	}
//	return false
//}
