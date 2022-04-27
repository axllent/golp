package cmd

import (
	"fmt"
	"os"

	"github.com/axllent/golp/app"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Build & watch src directories for changes",
	Long: `Build and watch your src directories for changes.
	
This will monitor your src directories for changes and instantly rebuild
their assets when a change is detected, useful for development.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if app.QuietLogging && app.VerboseLogging {
			fmt.Println("Error: cannot use --quiet and --verbose together")
			os.Exit(1)
		}

		if err := app.ParseConfig(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		app.WatchSrcDirs()
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.Flags().BoolVarP(&app.Minify, "minify", "m", false, "minify dist styles & scripts")
	watchCmd.Flags().StringVarP(&app.Conf.ConfigFile, "config", "c", "./golp.yaml", "config file")
	watchCmd.Flags().BoolVarP(&app.VerboseLogging, "verbose", "v", false, "verbose output")
	watchCmd.Flags().BoolVarP(&app.QuietLogging, "quiet", "q", false, "no output except for errors")
}
