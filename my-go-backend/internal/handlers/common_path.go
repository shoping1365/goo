package handlers

import (
	"os"
	"path/filepath"
)

// projectRoot returns the repository root directory (parent of my-go-backend).
// It is used by various handlers to resolve absolute paths under the public directory.
func projectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Dir(wd)
}
