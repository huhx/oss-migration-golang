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
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		current := viper.GetString("current")
		if current == setting.Local {
			path := viper.GetString("local.path")
			files, err := util.Lookup(path)
			if err == nil {
				resources := util.ResourceFilter(&files)
				markdownFiles := util.MarkdownFilter(&files)
				imagesInFile := util.ImageNames(&markdownFiles)

				abandonResources := lo.Filter(resources, func(item string, index int) bool {
					imageName := filepath.Base(item)
					markdownImage := util.FromMarkdown(imagesInFile, imageName)
					return markdownImage == nil
				})

				result := lo.Map(abandonResources, func(item string, index int) oss.PlanResponse {
					imageName := filepath.Base(item)

					return oss.PlanResponse{
						ImageName: imageName,
						Path:      item,
						ImageIn:   "local",
					}
				})

				if len(result) > 0 {
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader(util.GetStructFieldNames(reflect.TypeOf(oss.PlanResponse{})))

					for _, v := range result {
						table.Append([]string{
							v.ImageName,
							v.Path,
							v.ImageIn,
						})
					}
					table.Render()
				} else {
					fmt.Println("All the images are synced.")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
