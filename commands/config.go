package commands

// defaults read com.matryer.BitBar pluginsDirectory

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

// runMenu adds the `menu` subcommand to bitbar-openfaas
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runConfig,
}
var defaultGateway = "http://localhost:8080"
var defaultPrometheus = "http://localhost:9090"

// the questions to ask
var qs = []*survey.Question{
	{
		Name:   "gateway",
		Prompt: &survey.Input{Message: "Gateway URL?", Default: defaultGateway, Help: "URL of your OpenFaaS Gateway"},
	},
	{
		Name:   "prometheus",
		Prompt: &survey.Input{Message: "Prometheus URL?", Default: defaultPrometheus, Help: "URL of the OpenFaaS Prometheus instance"},
	},
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Nickname?", Help: "A short name to identify this OpenFaaS instance, for example \"Local OpenFaaS\""},
		Validate: survey.Required,
	},
}

func runConfig(cmd *cobra.Command, args []string) {
	// the answers will be written to this struct
	answers := struct {
		Name       string
		Gateway    string
		Prometheus string
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Results %+v.\n", answers)
	os.Exit(0)
}
