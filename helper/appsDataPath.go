package helper

import (
	"fmt"
	"os"
)

func AppsDataPath() string {
	appConfigDir, err := os.UserConfigDir()
	Fatal(fmt.Sprintf("Failed to get apps config directory: %s", err), true, err)

	return appConfigDir
}
