package main

import (
	"errors"

	codacy "github.com/codacy/codacy-engine-golang-seed"

	"os"
)

// GoReviveImplementation tool implementation
type GoReviveImplementation struct {
}

// Run runs the tool implementation
func (i GoReviveImplementation) Run(tool codacy.Tool, sourceDir string) ([]codacy.Issue, error) {
	configFile, _ := getConfigurationFile(tool.Patterns, sourceDir)
	if configFile != nil {
		defer os.Remove(configFile.Name())
	}

	filesToAnalyse, err := getListOfFilesToAnalyse(tool.Files, sourceDir)
	if err != nil {
		return nil, errors.New("Error getting files to analyse: " + err.Error())
	}

	reviveCmd := reviveCommand(configFile, filesToAnalyse, sourceDir)

	reviveOutput, err := reviveCmd.Output()
	if err != nil {
		return nil, errors.New("Error running revive: " + err.Error())
	}

	result := parseOutput(string(reviveOutput))

	return result, nil
}
