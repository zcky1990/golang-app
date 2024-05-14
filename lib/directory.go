package lib

import (
	"os"
	"path/filepath"
)

func FindRootDir(dir string) string {
	knownProjectItems := []string{"go.mod", "main.go"}
	for {
		for _, item := range knownProjectItems {
			if _, err := os.Stat(filepath.Join(dir, item)); err == nil {
				return dir
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}
