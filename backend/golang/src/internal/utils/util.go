package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func GetProjectRoot() (string, error) {
	_, exPath, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(exPath)

	for {
		parent := filepath.Dir(currentDir)
		if currentDir == parent {
			return "", errors.New("cannot get project root path: not found go.mod")
		}

		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		currentDir = parent
	}
}
