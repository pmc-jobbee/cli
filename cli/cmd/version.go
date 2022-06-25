package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pmc-jobbee/cli/cli/app"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print app version",
	Long:  `Print app version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s-%s (Built: %s)\n", app.AppName, app.AppVersion, app.AppGitHash, app.AppBuildDate)
	},
}
