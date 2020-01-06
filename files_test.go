package main

import (
	"fmt"
	"testing"
)

func TestGetListOfFilesToAnalyse(t *testing.T) {
	files := []string{"file1.go", "file2.go"}

	analysisFiles := getListOfFilesToAnalyse(files)

	if len(files) != len(analysisFiles) {
		fmt.Sprintf("Expected %s, got %s", files, analysisFiles)
	}
}
