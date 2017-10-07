package commands

import (
	"fmt"
	"os"

	"github.com/johnmccabe/bitbar"
	"github.com/johnmccabe/bitbar-openfaas/assets"
	"github.com/johnmccabe/bitbar-openfaas/config"
	"github.com/johnmccabe/bitbar-openfaas/faas"
	"github.com/johnmccabe/bitbar-openfaas/graphs"
	"github.com/johnmccabe/bitbar-openfaas/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
}

// runMenu adds the `menu` subcommand to bitbar-openfaas
var menuCmd = &cobra.Command{
	Use:     "menu",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runMenu,
}

func runMenu(cmd *cobra.Command, args []string) {
	ex, _ := os.Executable()

	cfg, err := config.Read("~/.openfaas/")
	if err != nil {
		plugin := bitbar.New()
		plugin.StatusLine("FNs: ❓\n")
		menu := plugin.NewSubMenu()
		if err.Error() == "Config file not found" {
			menu.Line("Initialise config").Bash(ex).Params([]string{"config"}).Terminal(true).Refresh(true)
		} else {
			menu.Line(fmt.Sprintf("Error: %v", err))
		}
		fmt.Print(plugin.Render())
		os.Exit(1)
	}

	functions := new([]types.Function)
	faas.GetFunctions(cfg[0].Gateway, functions)

	plugin := bitbar.New()
	plugin.StatusLine(fmt.Sprintf("FNs: %d | color=green\n", len(*functions)))
	menu := plugin.NewSubMenu()
	menu.Line("").Image(assets.Logo)
	menu.Line("---")

	menu.Line(cfg[0].Name).Font("Menlo-Bold")

	for _, function := range *functions {
		menu.Line(function.Name).Font("Menlo-Regular").Size(14)
		statsMenu := menu.NewSubMenu()
		statsMenu.Line("Actions").Font("Menlo-Bold")
		statsMenu.Line("Delete Function").Bash(ex).Params([]string{"delete", function.Name}).Refresh(true).Terminal(false)
		statsMenu.HR()
		statsMenu.Line("API Stats").Font("Menlo-Bold")
		statsMenu.Line(fmt.Sprintf("Current Replicas:			%d", function.Replicas))
		statsMenu.Line(fmt.Sprintf("Total Invocations:		%d", function.InvocationCount))
		statsMenu.HR()
		statsMenu.Line("Prometheus Stats").Font("Menlo-Bold")
		statsMenu.Line("Invocations per second")
		img, _ := graphs.FunctionGraph(cfg[0].Prometheus, function.Name, "rate(gateway_functions_seconds_count{function_name='%s'} [30s])")
		statsMenu.Line("").Image(img)
		statsMenu.Line("Aggregate Invocations")
		img, _ = graphs.FunctionGraph(cfg[0].Prometheus, function.Name, "gateway_function_invocation_total{function_name='%s'}")
		statsMenu.Line("").Image(img)
	}
	menu.Line("---")
	menu.Line("Refresh..").Refresh(true)

	fmt.Print(plugin.Render())
}
