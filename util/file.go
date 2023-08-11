package util

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"oss-migration/oss"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var ignoredFiles = []string{"README.md", ".DS_Store"}
var resourcesSuffix = []string{".svg", ".png", ".jpeg"}

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

func ResourceFilter(paths *[]string) []string {
	return lo.Filter(*paths, func(item string, i int) bool {
		extension := path.Ext(item)

		return lo.Contains(resourcesSuffix, extension)
	})
}

func MarkdownFilter(paths *[]string) []string {
	return lo.Filter(*paths, func(item string, i int) bool {
		extension := path.Ext(item)

		return extension == ".md"
	})
}

func ImageNames(paths *[]string) []oss.MarkdownImage {
	var images []oss.MarkdownImage

	for _, imagePath := range *paths {
		images = append(images, ExtractImageNames(imagePath)...)
	}
	return images
}

func FromMarkdown(images []oss.MarkdownImage, image string) *oss.MarkdownImage {
	for _, markdownImage := range images {
		if markdownImage.ImageName == image {
			return &markdownImage
		}
	}
	return nil
}

func ExtractImageNames(path string) []oss.MarkdownImage {
	imagePattern := regexp.MustCompile(`!\[(.*)]\(([^)]+)\)`)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error reading Markdown file:", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	var markdownImages []oss.MarkdownImage

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if matches := imagePattern.FindStringSubmatch(line); len(matches) > 0 {
			imagePath := matches[2]
			imageName := imagePath[strings.LastIndex(imagePath, "/")+1:]
			imageTag := matches[1]
			markdownImages = append(markdownImages, oss.MarkdownImage{
				ImageName:    imageName,
				MarkdownName: path,
				LineNumber:   lineNumber,
				ImageTag:     imageTag,
			})
		}

	}

	return markdownImages
}

func BytesToKBString(bytes int64) string {
	kb := float64(bytes) / 1024
	return strconv.FormatFloat(kb, 'f', 2, 64) + " KB"
}

func GetStructFieldNames(structType reflect.Type) []string {
	var fieldNames []string
	for i := 0; i < structType.NumField(); i++ {
		fieldNames = append(fieldNames, structType.Field(i).Name)
	}
	return fieldNames
}
