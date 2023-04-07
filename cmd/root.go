package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"jenrykster.github.io/cli_test/pkg/stringer"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli_test",
	Short: "A simple mocked CLI application",
	Long: `You can use this command to do exactly nothing. 
But you also have the flexibility to do even more nothing.`,
	Run: func(cmd *cobra.Command, args []string) {
		stringer.HandleRootCommand()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
