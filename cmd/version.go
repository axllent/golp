package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/axllent/golp/updater"
	"github.com/spf13/cobra"
)

var (
	// Version is the default application version, updated on release
	Version = "dev"

	// Repo on Github for updater
	Repo = "axllent/golp"

	// RepoBinaryName on Github for updater
	RepoBinaryName = "golp"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version & update information",
	Long: `Display the current version & update information (if available).
	
To view module versions, use the --modules flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		modules, _ := cmd.Flags().GetBool("modules")

		if modules {
			bi, ok := debug.ReadBuildInfo()
			if !ok {
				return errors.New("Failed to read build info")
			}

			fmt.Printf("%s %s is compiled with the following:\n\n", os.Args[0], Version)

			displayMap := map[string]bool{
				"github.com/bep/golibsass":       true,
				"github.com/evanw/esbuild":       true,
				"github.com/goreleaser/fileglob": true,
				"github.com/tdewolff/minify/v2":  true,
				"github.com/radovskyb/watcher":   true,
			}

			for _, dep := range bi.Deps {
				if _, ok := displayMap[dep.Path]; ok {
					fmt.Printf("%-32s %s\n", dep.Path, dep.Version)
				}
			}

			return nil
		}

		update, _ := cmd.Flags().GetBool("update")

		if update {
			return updateApp()
		}

		fmt.Printf("%s %s compiled with %s on %s/%s\n",
			os.Args[0], Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)

		latest, _, _, err := updater.GithubLatest(Repo, RepoBinaryName)
		if err == nil && updater.GreaterThan(latest, Version) {
			fmt.Printf(
				"\nUpdate available: %s\nRun `%s version -u` to update (requires read/write access to install directory).\n",
				latest,
				os.Args[0],
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().
		BoolP("update", "u", false, "update to latest version")

	versionCmd.Flags().
		BoolP("modules", "m", false, "display module versions")
}

func updateApp() error {
	rel, err := updater.GithubUpdate(Repo, RepoBinaryName, Version)
	if err != nil {
		return err
	}

	fmt.Printf("Updated %s to version %s\n", os.Args[0], rel)
	return nil
}
