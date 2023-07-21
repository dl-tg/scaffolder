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
	// Project name
	var name, yaml string
	// Initialize Git?
	var git bool

	var yamlPath string

	// Define and parse flags
	flag.StringVar(&name, "name", "", "Project name")
	flag.StringVar(&yaml, "yaml", "", "Config to use")
	flag.BoolVar(&git, "git", false, "Use git in project")
	flag.Parse()

	// If project name or path to yaml was not provied, print usage and exit with code 1
	if name == "" || yaml == "" {
		helper.Fatal("Usage: scaffold --name <projname> --yaml <path> --git? <boolean> (without angle brackets, ? - optional)", false)
	}

	// Initialize Git repository if git is true (user agreed)
	if git {
		helper.Git(name)
	}

	// Construct a path to the provided yaml config file
	if runtime.GOOS == "windows" {
		yamlPath = os.Getenv("USERPROFILE") + "\\scaffolder-configs\\" + yaml + ".yaml"
	} else {
		home, err := os.UserHomeDir()
		helper.Fatal(fmt.Sprintf("Failed to get home directory: %s", err), true, err)
		yamlPath = home + "/scaffolder-configs/" + yaml + ".yaml"
	}

	// Scaffold the directory structure
	utils.Scaffold(name, yamlPath)

}
