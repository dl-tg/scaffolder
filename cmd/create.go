/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"
	"scaffolder/pkg/helper"
	"scaffolder/pkg/utils"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the scaffolding",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		yaml, _ := cmd.Flags().GetString("yaml")
		name, _ := cmd.Flags().GetString("name")
		git, _ := cmd.Flags().GetBool("git")

		// Initialize Git repository if specified
		if git {
			helper.Git(name)
		}

		// Construct a path to the provided yaml config file
		var yamlPath string
		if runtime.GOOS == "windows" {
			yamlPath = os.Getenv("USERPROFILE") + "\\scaffolder-configs\\" + yaml + ".yaml"
		} else {
			home, err := os.UserHomeDir()
			helper.Fatal(fmt.Sprintf("Failed to get home directory: %s", err), true, err)
			yamlPath = home + "/scaffolder-configs/" + yaml + ".yaml"
		}

		// Scaffold the directory structure
		utils.Scaffold(name, yamlPath)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Set up the flags
	createCmd.Flags().String("name", "", "Project name")
	createCmd.Flags().String("yaml", "", "Config to use")
	createCmd.Flags().Bool("git", false, "Initialize Git repository")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("yaml")
}
