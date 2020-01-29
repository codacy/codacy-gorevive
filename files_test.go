package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetListOfFilesToAnalyse(t *testing.T) {
	files := []string{"file1.go", "file2.go"}

	analysisFiles := getListOfFilesToAnalyse(files, "./")

	assert.Equal(t, len(files), len(analysisFiles))
}
