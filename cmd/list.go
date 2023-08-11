package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"oss-migration/oss"
	"oss-migration/setting"
	"oss-migration/util"
	"path/filepath"
	"reflect"
	"strconv"
)

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the resources from local or remote",
	Run: func(cmd *cobra.Command, args []string) {
		current := viper.GetString("current")
		if current == setting.Local {
			path := viper.GetString("local.path")
			files, err := util.Lookup(path)
			if err == nil {
				resources := util.ResourceFilter(&files)
				markdownFiles := util.MarkdownFilter(&files)
				imagesInFile := util.ImageNames(&markdownFiles)

				result := lo.Map(resources, func(item string, index int) oss.ListResponse {
					imageName := filepath.Base(item)
					stat, _ := os.Stat(item)
					markdownImage := util.FromMarkdown(imagesInFile, imageName)

					return oss.ListResponse{
						ImageName:    imageName,
						ImagePath:    item,
						CreateTime:   stat.ModTime(),
						ImageSize:    stat.Size(),
						IsUsed:       markdownImage != nil,
						MarkdownName: &markdownImage.MarkdownName,
						LineNumber:   &markdownImage.LineNumber,
						ImageTag:     &markdownImage.ImageTag,
					}
				})

				if len(result) < 1 {
					fmt.Println("No data")
					return
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader(util.GetStructFieldNames(reflect.TypeOf(oss.ListResponse{})))

				for _, v := range result {
					table.Append([]string{
						v.ImageName,
						v.ImagePath,
						v.CreateTime.Format("2006-01-02 15:04:05"),
						util.BytesToKBString(v.ImageSize),
						strconv.FormatBool(v.IsUsed),
						lo.TernaryF(v.ImageTag == nil,
							func() string { return "" },
							func() string { return *v.ImageTag },
						),
						lo.TernaryF(v.LineNumber == nil,
							func() string { return "0" },
							func() string { return strconv.Itoa(*v.LineNumber) },
						),
						lo.TernaryF(v.MarkdownName == nil,
							func() string { return "" },
							func() string { return *v.MarkdownName },
						),
					})
				}
				table.Render()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
