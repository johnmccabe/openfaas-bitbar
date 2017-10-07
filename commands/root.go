package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// Execute initialises the root and all subcommands
func Execute() {
	rootCmd.Execute()
}

// rootCmd is the root command it runs the menuSubcommand by default
var rootCmd = &cobra.Command{
	Use:   "bitbar-openfaas",
	Short: "TODO",
	Long:  `TODO`,
	Run:   runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	ex, _ := os.Executable()
	stat, _ := os.Lstat(ex)
	if (stat.Mode() & os.ModeSymlink) != 0 {
		// Run menu subcommand by default when run via a symlink
		menuCmd.Run(cmd, args)
	} else {
		cmd.Help()
	}
}
