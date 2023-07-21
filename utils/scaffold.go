package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"scaffolder/helper"

	"gopkg.in/yaml.v3"
)

// Scaffold the directory structure
func Scaffold(name string, yamlpath string) {
	// Read and get YAML data
	yamlData, err := os.ReadFile(yamlpath)
	helper.Fatal(fmt.Sprintf("Failed to read YAML file: %s", err), true, err)

	// Create map for the directory structure
	var dirs map[string]map[string]string

	// Unmarshal the YAML into our map
	err = yaml.Unmarshal(yamlData, &dirs)
	helper.Fatal(fmt.Sprintf("Error unmarshalling YAML: %s", err), true, err)

	// Create project folder
	err = os.Mkdir(name, 0755)
	helper.Fatal(fmt.Sprintf("Error creating project folder: %s", err), true, err)

	// Navigate to the project folder
	err = os.Chdir(name)
	helper.Fatal(fmt.Sprintf("Failed to navigate to project folder: %s", err), true, err)

	// Scaffold the directory structure :: iterating over the map
	for folder, files := range dirs {
		// Create the folders and subdirectories if necessary
		err = os.MkdirAll(folder, 0755)
		helper.Fatal(fmt.Sprintf("Error creating folder %s: %v", folder, err), true, err)

		// Create the files :: iterating over files from the map and getting specified content
		for fileName, content := range files {
			// Construct a file path for the file
			filePath := filepath.Join(folder, fileName)
			// Create the directories before creating the file
			err = os.MkdirAll(filepath.Dir(filePath), 0755)
			helper.Fatal(fmt.Sprintf("Error creating directories for %s: %v", filePath, err), true, err)

			// Create the files at filePath path and add specified content
			err = os.WriteFile(filePath, []byte(content), 0644)
			helper.Fatal(fmt.Sprintf("Failed to create file %s: %s", fileName, err), true, err)
		}
	}
}
