package cmd

import (
	"os"

	"github.com/axllent/golp/app"
	"github.com/axllent/golp/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile & copy your assets (single run)",
	Long: `Compile & copy your assets in a single run.

By default SASS & JS files will include SourceMaps, which are used by browsers
to debug your code. Run with '-m' to disable SourceMaps and minify the output.`,
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"package"},
	Run: func(cmd *cobra.Command, args []string) {
		if app.QuietLogging && app.VerboseLogging {
			app.Log().Error("Cannot use --quiet and --verbose together")
			os.Exit(1)
		}

		if cmd.CalledAs() == "package" {
			app.Minify = true
		}

		if err := app.ParseConfig(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		sw := utils.StartTimer()

		if err := app.Clean(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		for _, task := range app.Conf.Tasks {
			if err := task.Process(""); err != nil {
				app.Log().Error(err.Error())
				os.Exit(1)
			}
		}

		app.Log().Infof("completed in %v", sw.Elapsed())
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&app.Minify, "minify", "m", false, "minify dist styles & scripts")
	buildCmd.Flags().StringVarP(&app.Conf.ConfigFile, "config", "c", "./golp.yaml", "config file")
	buildCmd.Flags().BoolVarP(&app.VerboseLogging, "verbose", "v", false, "verbose output")
	buildCmd.Flags().BoolVarP(&app.QuietLogging, "quiet", "q", false, "no output except for errors")
}
