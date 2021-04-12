package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main()  {
	filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path,d,err)
		return nil
	})
	files, _ := os.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}


	//files, _ := filepath.Glob("*")
	//fmt.Println(files)
}
