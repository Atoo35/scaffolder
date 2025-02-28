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

	var yamlPath string
	var configPath string

	// Define and parse command-line flags
	flag.StringVar(&name, "name", "", "Project name")
	flag.StringVar(&yaml, "yaml", "", "Config to use")
	flag.StringVar(&configPath, "configdir", "", "Path to custom config")
	flag.BoolVar(&git, "git", false, "Use git in project")
	flag.Parse()

	// If the project name or path to the YAML file was not provided, print usage and exit with code 1
	if name == "" || yaml == "" {
		helper.Fatal("Usage: scaffold --name <projname> --yaml <configname> --git? <boolean> (without angle brackets, ? - optional)", false)
	}

	// Initialize Git repository if 'git' flag is true (user agreed)
	if git {
		helper.Git(name)
	}

	// Check and set the path to the YAML config file
	if configPath == "" {
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
		// If a custom config path was provided, validate and use it
		if !helper.ValidateYamlPath(fmt.Sprintf("%s/%s.yaml", configPath, yaml), &yamlPath) {
			configPath = ""
		}
	}

	// Scaffold the directory structure using the provided project name and YAML config path
	utils.Scaffold(name, yamlPath)
}
