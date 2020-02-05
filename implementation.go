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
	configFile, err := getConfigurationFile(tool.Patterns, sourceDir)
	if err != nil {
		defer os.Remove(configFile.Name())
	}

	filesToAnalyse, err := getListOfFilesToAnalyse(tool.Files, sourceDir)
	if err != nil {
		return nil, errors.New("Error getting files to analyse: " + err.Error())
	}

	reviveCmd := reviveCommand(configFile, filesToAnalyse, sourceDir)

	reviveOutput, reviveError, err := runCommand(reviveCmd)
	if err != nil {
		return nil, errors.New("Error running revive: " + reviveError)
	}

	result := parseOutput(reviveOutput)
	return result, nil
}
