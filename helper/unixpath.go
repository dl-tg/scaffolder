package helper

import (
	"fmt"
	"os"
)

// UnixPath takes a YAML filename as input and returns the Unix path where the file should be located.
// It utilizes the user's home directory to construct the path.
// Parameters:
//   - yaml: The YAML filename (without the directory path).
//
// Returns:
//   - string: The Unix path formed by concatenating the user's home directory,
//     the ".scaffolder" directory, and the given YAML filename.
//
// Example:
//
//	yamlFilename := "config"
//	path := UnixPath(yamlFilename)
//	// Result: path will contain the Unix path to the "config.yaml" file within the .scaffolder folder in user's home directory.
func UnixPath(yaml string) string {
	// Get the user's home directory path and any potential error.
	home, err := os.UserHomeDir()
	Fatal(fmt.Sprintf("Failed to get home directory: %s", err), true, err)

	// Concatenate the home directory, ".scaffolder" directory, and the given YAML filename to form the Unix path.
	var unixpath string = home + "/.scaffolder/" + yaml + ".yaml"

	return unixpath
}
