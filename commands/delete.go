package commands

import (
	"fmt"
	"os"

	"github.com/johnmccabe/openfaas-bitbar/config"
	"github.com/johnmccabe/openfaas-bitbar/faas"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// runMenu adds the `menu` subcommand to openfaas-bitbar
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {

	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}

	if len(args) < 1 {
		os.Exit(1)
	}

	functionName := args[0]

	err = faas.DeleteFunction(cfg.Stacks[0].Gateway, functionName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
