package util

import (
	"os"
	"path/filepath"
	"strings"
)

func Lookup(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".vuepress") || info.Name() == "README.md" || info.Name() == ".DS_Store" {
			return nil
		}

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	})
	return files, err
}
