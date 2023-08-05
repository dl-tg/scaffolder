package main

import (
	"flag"

	"github.com/dl-tg/scaffolder/helper"
	"github.com/dl-tg/scaffolder/utils"
)

func main() {

	// Project name and YAML config filename
	var name, yaml string
	// Initialize Git?
	var git bool
	// Remember the path to custom config folder specified in configPath?
	var remember bool
	// Will be used to construct path to target YAML later
	var yamlPath string
	// Path to custom config folder
	var configPath string
	// Dictionary of custom variables
	var yamlVariables helper.YamlVariableMap = map[string]string{}

	// Define and parse command-line flags
	flag.StringVar(&name, "name", "", "Project name")
	flag.StringVar(&yaml, "yaml", "", "Config to use")
	flag.StringVar(&configPath, "configdir", "", "Path to custom config")
	flag.BoolVar(&git, "git", false, "Use git in project")
	flag.BoolVar(&remember, "remember", false, "Remember the config path")
	flag.Var(&yamlVariables, "variables", "Set variables to be used as comma seperated key value pairs eg key:value,key2:value2 ")
	flag.Parse()

	// If the project name or path to the YAML file was not provided, print usage and exit with code 1
	if yaml == "" {
		flag.Usage()
	}

	// Initialize Git repository if 'git' flag is true (user agreed)
	if git {
		helper.Git(name)
	}

	yamlPath = helper.GetYamlPath(configPath, yaml)

	if remember {
		helper.SaveConfigDir(yamlPath)
	}
	// Scaffold the directory structure using the provided project name and YAML config path
	utils.Scaffold(name, yamlPath, yamlVariables)
}
