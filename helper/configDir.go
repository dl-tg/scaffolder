package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get scaffolder default config path
func getConfigDirPath() string {
	var configDirPath string = filepath.Join(AppsDataPath(), "scaffolder")
	return configDirPath
}

// SaveConfigDir creates the scaffolder config directory and saves the provided config path 
func SaveConfigDir(configPath string) {
	configDirPath := getConfigDirPath()

	err := os.MkdirAll(configDirPath, 0755)
	Fatal(fmt.Sprintf("Error: Failed to create scaffolder config folder: %s", err), true, err)

	configFilePath := filepath.Join(configDirPath, "configDir.txt")
	f, err := os.Create(configFilePath)
	Fatal(fmt.Sprintf("Error: Failed to create configDir text file in scaffolder folder: %s", err), true, err)

	defer f.Close()
	_, err = f.WriteString(configPath)
	Fatal(fmt.Sprintf("Error: Failed to write custom config path to configDir.txt: %s", err), true, err)
}

// GetConfigDir retrieves the saved config path from the scaffolder config directory
func GetConfigDir() string {
	configFilePath := filepath.Join(getConfigDirPath(), "configDir.txt")
	f, err := os.OpenFile(configFilePath, os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("Error: Could not open configDir text file: %s", err)
	}

	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	return string(buf[:n])
}
