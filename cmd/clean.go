package cmd

import (
	"os"

	"github.com/axllent/golp/app"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean (delete) your configured files/directories",
	Long:  `Clean (delete) your configured files/directories.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.ParseConfig(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		if err := app.Clean(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
