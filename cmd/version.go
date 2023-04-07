package cmd

import (
	"github.com/spf13/cobra"
	"jenrykster.github.io/cli_test/pkg/stringer"
)

func init() {

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Just the software version. Nothing strange about it.",
	Run: func(cmd *cobra.Command, args []string) {
		stringer.HandleVersionCommand()
	},
}
