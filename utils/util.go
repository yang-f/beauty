package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func CurrentPath() string {
    return getCurrentDirectory()
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
