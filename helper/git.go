package helper

import (
	"fmt"
	"os"
	"os/exec"
)

// Git initializes a Git repository for the specified project.
// It navigates to the project directory and runs the "git init" command to create a new Git repository.
// If any error occurs during the process, it logs a fatal error message and terminates the program.
//
// Parameters:
//   - name: The name of the project (directory) where the Git repository will be initialized.
//
// Example:
//
//	projectName := "my_project"
//	Git(projectName)
//	// Result: A new Git repository will be created in the "my_project" directory.
func Git(name string) {
	// Navigate to project directory
	os.Chdir(name)

	// Initialize git repository
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	Fatal(fmt.Sprintf("Error initializing Git repository: %s", err), true, err)

	fmt.Println("Git repository initialized!")
}
