package create

import (
	"os"
	"path/filepath"

	jww "github.com/spf13/jwalterweatherman"
)

func NewContent(targetPath string) error {
	path_ := filepath.Join("content", targetPath)
	//if isForce {
	dirPath := filepath.Dir(path_)
	err := os.MkdirAll(dirPath, 0766)
	if err != nil {
		return err
	}
	//}

	jww.WARN.Printf("attempting to create %q", targetPath)

	if filepath.Ext(path_) == "" {
		path_ += ".gmi"
		jww.WARN.Printf("auto add ext .gmi")
	}

	f, err := os.Create(path_)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
