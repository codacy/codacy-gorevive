package main

import (
	"errors"

	codacy "github.com/codacy/codacy-golang-tools-engine"

	"os"
)

var formatter = "ndjson"

// GoReviveImplementation tool implementation
type GoReviveImplementation struct {
}

// Run runs the tool implementation
func (i GoReviveImplementation) Run(tool codacy.Tool, sourceDir string) ([]codacy.Issue, error) {
	patternToToml(codacy.Pattern{})
	configFile, _ := getConfigurationFile(tool.Patterns, sourceDir)
	if configFile != nil {
		defer os.Remove(configFile.Name())
	}

	filesToAnalyse := getListOfFilesToAnalyse(tool.Files, sourceDir)

	reviveCmd := command(configFile, filesToAnalyse, sourceDir)

	reviveOutput, err := reviveCmd.Output()
	if err != nil {
		return nil, errors.New("Error running revive: " + err.Error())
	}

	result := parseOutput(string(reviveOutput))

	return result, nil
}
