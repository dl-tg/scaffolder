package main

import (
	"flag"
	"scaffolder/helper"
	"scaffolder/utils"
)

func main() {
	// Project name and path to YAML
	var name, yaml string
	// Initialize Git?
	var git bool

	// Define and parse flags
	flag.StringVar(&name, "name", "", "Project name")
	flag.StringVar(&yaml, "yaml", "", "Path to YAML with directory structure")
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

	// Scaffold the directory structure
	utils.Scaffold(name, yaml)

}
