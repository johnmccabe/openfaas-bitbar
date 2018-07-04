package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/johnmccabe/bitbar"
	"github.com/johnmccabe/openfaas-bitbar/assets"
	"github.com/johnmccabe/openfaas-bitbar/config"
	"github.com/johnmccabe/openfaas-bitbar/faas"
	"github.com/johnmccabe/openfaas-bitbar/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
}

var menuCmd = &cobra.Command{
	Use:    "menu",
	Run:    runMenu,
	Hidden: true,
}

func runMenu(cmd *cobra.Command, args []string) {
	ex, _ := os.Executable()
	cfg, err := config.Read()
	if err != nil {
		plugin := bitbar.New()
		plugin.StatusLine(" ❓").Font("Avenir").Size(16).TemplateImage(assets.MonoLogo)
		menu := plugin.NewSubMenu()
		if err.Error() == "Config file not found" {
			menu.Line("Initialise config").Bash(ex).Params([]string{"config"}).Terminal(true).Refresh(true)
		} else {
			menu.Line(fmt.Sprintf("Error: %v", err))
		}
		fmt.Print(plugin.Render())
		os.Exit(1)
	}

	totalFnCount := 0

	functionMap := map[int][]types.Function{}

	for authIndex, auth := range cfg.Auths {
		functions := []types.Function{}
		faas.GetFunctions(auth, &functions)
		functionMap[authIndex] = functions
		totalFnCount = totalFnCount + len(functions)
	}

	menubarIcon := themeAppropriateIcon()

	plugin := bitbar.New()
	plugin.StatusLine(fmt.Sprintf(" %d", totalFnCount)).Font("Avenir").Size(16).Image(menubarIcon)

	menu := plugin.NewSubMenu()
	menu.Line("").Image(assets.Logo)

	for authIndex, functions := range functionMap {

		menu.Line("---")

		menu.Line(fmt.Sprintf("GW: %s", cfg.Auths[authIndex].Gateway)).Font("Menlo-Bold").Bash("/usr/bin/open").Params([]string{cfg.Auths[authIndex].Gateway}).Terminal(false)

		for _, function := range functions {
			menu.Line(fmt.Sprintf("⁃ %s", function.Name)).Font("Menlo-Regular").Size(14)
			statsMenu := menu.NewSubMenu()
			statsMenu.Line("Actions").Font("Menlo-Bold").Color("black")
			statsMenu.Line("Delete Function").Bash(ex).Params([]string{"delete", function.Name, strconv.Itoa(authIndex)}).Refresh(true).Terminal(false)
			statsMenu.HR()
			statsMenu.Line("Function Stats").Font("Menlo-Bold").Color("black")
			statsMenu.Line(fmt.Sprintf("Current Replicas:			%d", function.Replicas)).Color("black")
			statsMenu.Line(fmt.Sprintf("Available Replicas:		%d", function.AvailableReplicas)).Color("black")
			statsMenu.Line(fmt.Sprintf("Total Invocations:		%d", function.InvocationCount)).Color("black")
			statsMenu.HR()
			statsMenu.Line("Image").Font("Menlo-Bold").Color("black")
			statsMenu.Line(function.Image).Color("black")
			statsMenu.HR()
			if len(function.EnvProcess) > 0 {
				statsMenu.Line("EnvProcess").Font("Menlo-Bold").Color("black")
				statsMenu.Line(function.EnvProcess).Color("black")
				statsMenu.HR()
			}
			if len(function.Labels) > 0 {
				statsMenu.Line("Labels").Font("Menlo-Bold").Color("black")
				for k, v := range function.Labels {
					statsMenu.Line(fmt.Sprintf("%s=%s", k, v)).Color("black")
				}
			}
		}
	}

	menu.Line("---")
	menu.Line("Refresh..").Refresh(true)

	fmt.Print(plugin.Render())
}

func themeAppropriateIcon() string {
	if _, ok := os.LookupEnv("BitBarDarkMode"); ok {
		return assets.WhiteLogo
	}
	return assets.MonoLogo
}
