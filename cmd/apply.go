package cmd

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"oss-migration/setting"
	"oss-migration/util"
	"path/filepath"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
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

				if len(abandonResources) < 1 {
					fmt.Println("All the data are synced.")
					return
				}

				for _, resource := range abandonResources {
					err := os.Remove(resource)
					if err == nil {
						fmt.Printf("%s in %s is removed.\n", filepath.Base(resource), resource)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
