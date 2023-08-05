package helper

import (
	"fmt"
	"os"
)

// ValidateYamlPath checks if the given path exists in the filesystem.
// If the path does not exist, it returns false.
// If any other error occurs, it logs a fatal error message and returns false.
// Parameters:
//   - path: The path to validate.
//
// Returns:
//   - bool: True if the path exists and is valid, false otherwise.
func ValidateYamlPath(path string) bool {
	// Check if the path exists in the filesystem.
	_, err := os.Stat(path)
	if err != nil {
		// If the path does not exist, return false.
		if os.IsNotExist(err) {
			return false
		} else {
			// If any other error occurs, log a fatal error message and return false.
			Fatal(fmt.Sprintf("Error occurred while checking absolute path: %s", err), true, err)
			return false
		}
	}

	return true
}
