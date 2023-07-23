package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"scaffolder/helper"
	"scaffolder/utils"
)

func main() {

	// Project name and YAML config filename
	var name, yaml string
	// Initialize Git?
	var git bool
	var remember bool
	var yamlPath string
	var configPath string
	var yamlVariables yamlVariableMap = map[string]string{}

	// Define and parse command-line flags
	flag.StringVar(&name, "name", "", "Project name")
	flag.StringVar(&yaml, "yaml", "", "Config to use")
	flag.StringVar(&configPath, "configdir", "", "Path to custom config")
	flag.BoolVar(&git, "git", false, "Use git in project")
	flag.BoolVar(&remember, "remember", false, "Remember the config path")
	flag.Var(&yamlVariables, "values", "key value pairs")
	flag.Parse()
	fmt.Println(yamlVariables["age"])

	// If the project name or path to the YAML file was not provided, print usage and exit with code 1
	if name == "" || yaml == "" {
		helper.Fatal("Usage: scaffold --name <projname> --yaml <configname> --configdir? <custom config path> --git? <boolean> --remember? <boolean> (without angle brackets, ? - optional) ", false)
	}

	// Initialize Git repository if 'git' flag is true (user agreed)
	if git {
		helper.Git(name)
	}

	// Check and set the path to the YAML config file
	if configPath == "" {
		savedPath, err := helper.GetConfigDir()
		helper.Fatal(fmt.Sprintf("Could not get config path: %s", err), true, err)

		if savedPath == "" {
			// Construct default paths for the YAML file based on the user's operating system
			var unixDefaultPath string = helper.UnixPath(yaml)
			var winDefaultPath string = os.Getenv("USERPROFILE") + "\\.scaffolder\\" + yaml + ".yaml"
			var defaultPathExists bool

			if runtime.GOOS == "windows" {
				defaultPathExists = helper.ValidateYamlPath(winDefaultPath, &yamlPath)
			} else {
				defaultPathExists = helper.ValidateYamlPath(unixDefaultPath, &yamlPath)
			}

			// If the default path does not exist, try the YAML file in the current directory
			if !defaultPathExists {
				if !helper.ValidateYamlPath(fmt.Sprintf("./%s.yaml", yaml), &yamlPath) {
					helper.Fatal(fmt.Sprintf("Could not find %s.yaml", yaml), false)
				}
			}
		} else {
			yamlPath = savedPath
		}

	} else {
		// If a custom config path was provided, validate and use it
		if !helper.ValidateYamlPath(fmt.Sprintf("%s/%s.yaml", configPath, yaml), &yamlPath) {
			configPath = ""
			// Store the config path in the database
		}
	}

	// Store the path in the database
	if remember {
		err := helper.SaveConfigDir(yamlPath)
		helper.Fatal(fmt.Sprintf("Could not save config path: %s", err), true, err)
	}
	// Scaffold the directory structure using the provided project name and YAML config path
	utils.Scaffold(name, yamlPath, *&yamlVariables)
}
