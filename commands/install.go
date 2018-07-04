package commands

// defaults read com.matryer.BitBar pluginsDirectory

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "TODO",
	Example: `TODO`,
	Long:    `TODO`,
	Run:     runInstall,
}

var defaultRefreshRate = "30s"

func runInstall(cmd *cobra.Command, args []string) {
	ex, _ := os.Executable()

	installQuestions := []*survey.Question{
		{
			Name:   "refreshRate",
			Prompt: &survey.Input{Message: "Refresh rate?", Default: defaultRefreshRate, Help: "Plugin refresh rate, for example: 30s, 1m, 5m"},
		},
	}

	// the answers will be written to this struct
	answers := struct {
		RefreshRate string
	}{}

	// perform the questions
	err := survey.Ask(installQuestions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	pluginDir, err := getPluginDir()
	if err != nil {
		fmt.Printf("Did you install bitbar? %v", err)
		os.Exit(1)
	}

	pluginSymlink := fmt.Sprintf("%s/openfaas-bitbar.%s", pluginDir, answers.RefreshRate)

	err = os.Symlink(ex, pluginSymlink)
	if err != nil {
		fmt.Printf("Unable to create symlink: %s for binary: %s", pluginSymlink, ex)
		os.Exit(1)
	}

}

func getPluginDir() (string, error) {
	cmdName := "/usr/bin/defaults"
	cmdArgs := []string{"read", "com.matryer.BitBar", "pluginsDirectory"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("unable to determine pluginsDirectory: %v, %s", err, string(cmdOut))
	}

	dir := strings.TrimRight(string(cmdOut), "\n")

	if !dirExists(dir) {
		return "", fmt.Errorf("unable to check if dir exists: %v, %s", err, dir)
	}

	return dir, nil
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true
	}
	return false
}
