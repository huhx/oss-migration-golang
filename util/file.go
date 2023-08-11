package util

import (
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
)

var ignoredFiles = []string{"README.md", ".DS_Store"}

func Lookup(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".vuepress") || lo.Contains(ignoredFiles, info.Name()) {
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
