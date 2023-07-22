package helper

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Save the configpath to a local file called configDir.txt in the user's home directory
func SaveConfigDir(configPath string) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(home, "configDir.txt"))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(configPath)
	if err != nil {
		return err
	}
	return nil
}

// Get the configpath from the local file called configDir.txt in the user's home directory
func GetConfigDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	f, err := os.OpenFile(filepath.Join(home, "configDir.txt"), os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	return string(buf[:n]), nil
}
