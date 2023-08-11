package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"oss-migration/setting"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:       "current",
	Short:     "Get current choose for storing resources",
	ValidArgs: []string{"local", "remote"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			current := viper.GetString("current")
			if current == setting.Local {
				fmt.Printf("Current: %s, and path is %s\n", current, viper.GetString("local.path"))
			} else {
				fmt.Printf("Current: %s, and endpoint is %s\n", current, viper.GetString("remote.endpoint"))
			}
		} else {
			viper.Set("current", args[0])
			_ = viper.WriteConfig()
		}
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
