package util

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"oss-migration/oss"
	"path"
	"path/filepath"
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

func FromMarkdown(images []oss.MarkdownImage, image string) *string {
	for _, markdownImage := range images {
		if markdownImage.ImageName == image {
			return &markdownImage.MarkdownName
		}
	}
	return nil
}

func ExtractImageNames(path string) []oss.MarkdownImage {
	markdownContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading Markdown file:", err)
		return nil
	}

	imageTagPattern := regexp.MustCompile(`!\[.*]\(([^)]+)\)`)
	imageTags := imageTagPattern.FindAllStringSubmatch(string(markdownContent), -1)

	var markdownImages []oss.MarkdownImage
	for _, tag := range imageTags {
		if len(tag) > 1 {
			imagePath := tag[1]
			imageName := imagePath[strings.LastIndex(imagePath, "/")+1:]
			markdownImages = append(markdownImages, oss.MarkdownImage{
				ImageName:    imageName,
				MarkdownName: path,
			})
		}
	}
	return markdownImages
}

func BytesToKBString(bytes int64) string {
	kb := float64(bytes) / 1024
	return strconv.FormatFloat(kb, 'f', 2, 64) + " KB"
}
