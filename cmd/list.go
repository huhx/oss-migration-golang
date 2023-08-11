package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"oss-migration/setting"
	"oss-migration/util"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the resources from local or remote",
	Run: func(cmd *cobra.Command, args []string) {
		current := viper.GetString("current")
		if current == setting.Local {
			path := viper.GetString("local.path")
			files, err := util.Lookup(path)
			if err == nil {
				fmt.Println(files)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
