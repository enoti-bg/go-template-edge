package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "{{cookiecutter.project_name}}",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error at launch: [%s]", err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config-path", "./resources/config", "Configuration directory for the environment dependent files.")
	rootCmd.MarkFlagRequired("config-path")
	cobra.OnInitialize()
}
