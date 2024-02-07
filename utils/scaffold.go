package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dl-tg/scaffolder/helper"

	"gopkg.in/yaml.v3"
)

// Pattern to find all appearance of curly braces enclosed instances
var variableMatchPattern = regexp.MustCompile(`\{[^{}]+\}`)

// replaceVariables takes a string content and a map containing set variables with their values
// the string that is returned has all occurences of a variable refrence replaced with the actual value that was set
func replaceVariables(name string, variableMap map[string]string) string {
	matches := variableMatchPattern.FindAllString(name, -1)
	// Check if there is at least a pattern match
	if len(matches) <= 0 {
		return name
	}
	// Loop through collected variables set and replace any occurence
	for k, v := range variableMap {
		// Construct regex pattern using our varibale key
		pattern := fmt.Sprintf(`\{%s\}`, regexp.QuoteMeta(k))
		// Compile the regular expression
		regex := regexp.MustCompile(pattern)
		name = regex.ReplaceAllString(name, v)
	}
	return name
}

/*
Scaffold creates a directory structure based on the provided YAML file.
It reads the YAML file at the given 'yamlpath', unmarshals its contents,
and creates the specified folders and files accordingly in the current directory.

Parameters:

  - name: The name of the project folder to be created.

  - yamlpath: The path to the YAML file containing the directory structure information.

Example Usage:

	projectName := "my_project"
	yamlFilePath := "/path/to/structure.yaml"
	Scaffold(projectName, yamlFilePath)

	// Result: The directory structure defined in the "structure.yaml" file will be created in the "my_project" directory.

Note:

  - The function assumes that the 'yamlpath' points to a valid YAML file that adheres to the expected format.

  - It will create the necessary folders and subdirectories as specified in the YAML file.

  - If any error occurs during the process, the function will log a fatal error message and exit the program.
*/
func Scaffold(name string, yamlpath string, setVariables map[string]string) {
	// Read and get YAML data
	yamlData, err := os.ReadFile(yamlpath)
	helper.Fatal(fmt.Sprintf("Failed to read YAML file: %s", err), true, err)

	// Create map for the directory structure
	var dirs map[string]interface{}

	// Unmarshal the YAML into our map
	err = yaml.Unmarshal(yamlData, &dirs)
	helper.Fatal(fmt.Sprintf("Error unmarshalling YAML: %s", err), true, err)

	// Create project folder if name was specified, else scaffold in current directoy
	if name != "" {
		err := os.Mkdir(name, 0755)
		helper.Fatal(fmt.Sprintf("Error creating project folder: %s", err), true, err)
		// Navigate to the project folder
		err = os.Chdir(name)
		helper.Fatal(fmt.Sprintf("Failed to navigate to project folder: %s", err), true, err)
	}

	// Scaffold the directory structure :: iterating over the map
	scaffoldDirs(".", dirs, setVariables)
}

func scaffoldDirs(basePath string, dirs map[string]interface{}, setVariables map[string]string) {
	for name, item := range dirs {
		// Apply variable replacement to the name
		name = replaceVariables(name, setVariables)
		switch v := item.(type) {
		case string: // If it's a file, create the file
			createFile(basePath, name, v, setVariables)
		case map[string]interface{}: // If it's a directory, recursively call scaffoldDirs
			newPath := createDirectory(basePath, name, v, setVariables)
			scaffoldDirs(newPath, v, setVariables)
		default:
			if strings.Contains(name, ".") {
				createFile(basePath, name, "", setVariables)
			} else {
				createDirectory(basePath, name, nil, setVariables)
			}
		}
	}
}

func createFile(basePath, name, content string, setVariables map[string]string) {
	filePath := filepath.Join(basePath, name)
	content = replaceVariables(content, setVariables)
	err := os.WriteFile(filePath, []byte(content), 0644)
	helper.Fatal(fmt.Sprintf("Failed to create file %s: %s", filePath, err), true, err)
}

func createDirectory(basePath, name string, content map[string]interface{}, setVariables map[string]string) string {
	newPath := filepath.Join(basePath, name)
	err := os.Mkdir(newPath, 0755)
	helper.Fatal(fmt.Sprintf("Error creating folder %s: %v", newPath, err), true, err)
	return newPath
}
