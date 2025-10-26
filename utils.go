package main

import (
	"os"
	"path/filepath"
)

func expandPath(path string) string {
	if len(path) > 0 && path[:2] == "~/" {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}
