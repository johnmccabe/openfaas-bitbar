package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/johnmccabe/openfaas-bitbar/config"
	"github.com/johnmccabe/openfaas-bitbar/faas"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:    "delete",
	Run:    runDelete,
	Hidden: true,
}

func runDelete(cmd *cobra.Command, args []string) {

	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}

	if len(args) < 2 {
		os.Exit(1)
	}

	functionName := args[0]
	authIndex, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = faas.DeleteFunction(cfg.Auths[authIndex], functionName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
