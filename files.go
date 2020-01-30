package main

import (
	"os"
	"path/filepath"
	"strings"
)

func getListOfFilesToAnalyse(files []string, sourceDir string) ([]string, error) {
	if len(files) > 0 {
		return files, nil
	}

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}
