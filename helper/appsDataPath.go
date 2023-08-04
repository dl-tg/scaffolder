package helper

import (
	"fmt"
	"os"
)

// AppsDataPath returns the path to the user-specific configuration directory for apps.
// It retrieves the path using os.UserConfigDir() function and returns it as a string.
// If an error occurs while getting the path, it will print the error message and exit the program.
func AppsDataPath() string {
	appConfigDir, err := os.UserConfigDir()
	Fatal(fmt.Sprintf("Failed to get apps config directory: %s", err), true, err)

	return appConfigDir
}
