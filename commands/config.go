package commands

// defaults read com.matryer.BitBar pluginsDirectory

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/johnmccabe/openfaas-bitbar/config"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

// runMenu adds the `menu` subcommand to openfaas-bitbar
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runConfig,
}

func runConfig(cmd *cobra.Command, args []string) {
	cfg, _ := config.Read()
	newCfg := new(config.Config)

	if len(cfg.Stacks) > 0 {
		for _, stack := range cfg.Stacks {
			stackResponse, err := promptForStack(stack)
			if err != nil {
				fmt.Printf("Error gathering stack details: [%v]", err)
				os.Exit(1)
			}
			newCfg.Stacks = append(newCfg.Stacks, stackResponse)
		}

	} else {
		stackResponse, err := promptForStack(config.Stack{})
		if err != nil {
			fmt.Printf("Error gathering stack details: [%v]", err)
			os.Exit(1)
		}
		newCfg.Stacks = append(newCfg.Stacks, stackResponse)
	}

	config.EnsureConfigDir()

	f, err := os.OpenFile(config.File(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Unable to expand homedir: [%v]", err)
		os.Exit(1)
	}

	defer f.Close()

	y, _ := yaml.Marshal(newCfg)

	if _, err = f.Write(y); err != nil {
		fmt.Printf("Unable to save config: [%v]", err)
		os.Exit(1)
	}

	fmt.Println("Config successfully updated.")

	refreshPlugin()

	os.Exit(0)
}

func refreshPlugin() error {
	answers := struct {
		Refresh bool
	}{}

	var refreshQuestions = []*survey.Question{
		{
			Name: "refresh",
			Prompt: &survey.Confirm{
				Message: "Refresh plugin (window will lose focus)?",
			},
		},
	}

	err := survey.Ask(refreshQuestions, &answers)
	if err != nil {
		return err
	}

	if answers.Refresh {
		cmd := exec.Command("open", "bitbar://refreshPlugin?name=openfaas-bitbar.*?")
		fmt.Println("Refreshing plugin...")
		err := cmd.Run()
		if err != nil {
			log.Printf("Command exited unexpectedly with error: %v", err)
		}

	}

	return nil
}

func promptForStack(stack config.Stack) (config.Stack, error) {
	answers := struct {
		Name       string
		Gateway    string
		Prometheus string
	}{}

	err := survey.Ask(stackQuestions(stack), &answers)
	if err != nil {
		return config.Stack{}, err
	}

	stack = config.Stack{
		Name:       answers.Name,
		Gateway:    answers.Gateway,
		Prometheus: answers.Prometheus,
	}

	return stack, nil
}

func stackQuestions(stack config.Stack) []*survey.Question {

	if (config.Stack{}) == stack {
		stack = config.DefaultStack
	}

	stackQuestions := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Nickname?",
				Default: stack.Name,
				Help:    "A short name to identify this OpenFaaS instance, for example \"Local OpenFaaS\"",
			},
			Validate: survey.Required,
		},
		{
			Name: "gateway",
			Prompt: &survey.Input{
				Message: "Gateway URL?",
				Default: stack.Gateway,
				Help:    "URL of your OpenFaaS Gateway"},
			Validate: survey.Required,
		},
		{
			Name: "prometheus",
			Prompt: &survey.Input{
				Message: "Prometheus URL?",
				Default: stack.Prometheus,
				Help:    "URL of the OpenFaaS Prometheus instance",
			},
			Validate: survey.Required,
		},
	}
	return stackQuestions
}

func editStacks(stackNames []string) []*survey.Question {
	editStackQuestions := []*survey.Question{
		{
			Name: "menuChoice",
			Prompt: &survey.Select{
				Message: "Edit existing stack?",
				Options: stackNames,
			},
		},
	}
	return editStackQuestions
}
