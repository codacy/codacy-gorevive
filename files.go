package main

import (
	"path/filepath"
)

func getListOfFilesToAnalyse(files []string, sourceDir string) []string {
	if len(files) > 0 {
		return files
	}

	globSubFolder := filepath.Join(sourceDir, "**/*.go")
	globRootFolder := filepath.Join(sourceDir, "*.go")
	filesSub, _ := filepath.Glob(globSubFolder)
	filesRoot, _ := filepath.Glob(globRootFolder)
	return append(filesSub, filesRoot...)
}
