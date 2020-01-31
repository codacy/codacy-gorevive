package main

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

const resourceFolder = "files_test"

func TestGetListOfFilesToAnalyse(t *testing.T) {
	files := []string{"file1.go", "file2.go"}

	analysisFiles, err := getListOfFilesToAnalyse(files, path.Join(testsResourcesFolder, resourceFolder))

	assert.Nil(t, err)
	assert.Equal(t, len(files), len(analysisFiles))
}

func TestGetListOfFilesToAnalyseWithEmptyFilesList(t *testing.T) {
	files := []string{}
	expectedFiles := []string{"test/data/files_test/bar/test/example2.go", "test/data/files_test/example.go"}
	analysisFiles, err := getListOfFilesToAnalyse(files, path.Join(testsResourcesFolder, resourceFolder))

	assert.Nil(t, err)
	assert.Equal(t, len(expectedFiles), len(analysisFiles))
	assert.Equal(t, expectedFiles, analysisFiles)
}
