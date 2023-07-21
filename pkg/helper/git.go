package helper

import (
	"fmt"
	"os"
	"os/exec"
)

// Initialize Git repository for specified project
func Git(name string) {
	// Navigate to project directory
	os.Chdir(name)

	// Initialize git repository
	cmd := exec.Command("git", "init")
	err := cmd.Run()
	Fatal(fmt.Sprintf("Error initializing Git repository: %s", err), true, err)

	fmt.Println("Git repository initialized!")
}
