package main

import (
	"path/filepath"
)

func getListOfFilesToAnalyse(files []string) []string {
	if len(files) > 0 {
		return files
	}

	files, _ = filepath.Glob("**/*.go")
	return files
}
