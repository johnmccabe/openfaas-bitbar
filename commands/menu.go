package commands

import (
	"fmt"
	"os"

	"github.com/johnmccabe/bitbar"
	"github.com/johnmccabe/openfaas-bitbar/assets"
	"github.com/johnmccabe/openfaas-bitbar/config"
	"github.com/johnmccabe/openfaas-bitbar/faas"
	"github.com/johnmccabe/openfaas-bitbar/graphs"
	"github.com/johnmccabe/openfaas-bitbar/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
}

// runMenu adds the `menu` subcommand to openfaas-bitbar
var menuCmd = &cobra.Command{
	Use:     "menu",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runMenu,
}

func runMenu(cmd *cobra.Command, args []string) {
	ex, _ := os.Executable()

	if config.LegacyTOMLExists() {
		plugin := bitbar.New()
		plugin.StatusLine("").Image(assets.RedLogo)
		menu := plugin.NewSubMenu()
		menu.Line("TOML config detected: ~/.openfaas/config.toml")
		menu.Line("Remove it before running `openfaas-bitbar config`.")
		fmt.Print(plugin.Render())
		os.Exit(1)
	}

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

	functions := new([]types.Function)
	faas.GetFunctions(cfg.Stacks[0].Gateway, functions)

	menubarIcon := themeAppropriateIcon()

	plugin := bitbar.New()
	plugin.StatusLine(fmt.Sprintf(" %d", len(*functions))).Font("Avenir").Size(16).Image(menubarIcon)
	menu := plugin.NewSubMenu()
	menu.Line("").Image(assets.Logo)
	menu.Line("---")

	menu.Line(cfg.Stacks[0].Name).Font("Menlo-Bold")

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
		img, _ := graphs.FunctionGraph(cfg.Stacks[0].Prometheus, function.Name, "rate(gateway_functions_seconds_count{function_name='%s'} [30s])")
		statsMenu.Line("").Image(img)
		statsMenu.Line("Aggregate Invocations")
		img, _ = graphs.FunctionGraph(cfg.Stacks[0].Prometheus, function.Name, "gateway_function_invocation_total{function_name='%s'}")
		statsMenu.Line("").Image(img)
	}

	menu.Line("---")
	menu.Line("Open Gateway...").Bash("/usr/bin/open").Params([]string{cfg.Stacks[0].Gateway}).Terminal(false)
	menu.Line("Open Prometheus...").Bash("/usr/bin/open").Params([]string{cfg.Stacks[0].Prometheus}).Terminal(false)

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
