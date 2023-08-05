package helper

import (
	"fmt"
	"path/filepath"
)

func GetYamlPath(configPath string, yaml string) string {
	var pathToUse string = ""
	var defaultPath string = filepath.Join(AppsDataPath(), "scaffolder", yaml+".yaml")
	var savedPath string = GetConfigDir()
	var customPath string = fmt.Sprintf("%s/%s.yaml", configPath, yaml)
	var routePath string = fmt.Sprintf("./%s", yaml)

	// Set the path to the YAML file based on whether the user specified a custom config path or not. If not, a saved or a default file will be used
	if configPath == "" {
		if savedPath == "" {
			if defaultPath == "" {
				pathToUse = defaultPath
			} else {
				pathToUse = routePath
			}
		} else {
			pathToUse = savedPath
		}
	} else {
		pathToUse = customPath
	}

	// Check if the YAML file exists
	if !ValidateYamlPath(pathToUse) {
		Fatal(fmt.Sprintf("Could not find %s", pathToUse), false)
	}

	return pathToUse
}
