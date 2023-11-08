package main

import (
	"errors"
	"context"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"

	"os"
)

// GoReviveImplementation tool implementation
type GoReviveImplementation struct {	
}

// Run runs the tool implementation
func (i GoReviveImplementation) Run(ctx context.Context, tool codacy.ToolExecution) ([]codacy.Result, error) {
	configFile, err := getConfigurationFile(*tool.Patterns, tool.SourceDir)
	if err == nil {
		defer os.Remove(configFile.Name())
	}

	filesToAnalyse, err := getListOfFilesToAnalyse(*tool.Files, tool.SourceDir)
	if err != nil {
		return nil, errors.New("Error getting files to analyse: " + err.Error())
	}

	reviveCmd := reviveCommand(configFile, filesToAnalyse, tool.SourceDir)

	reviveOutput, reviveError, err := runCommand(reviveCmd)
	if err != nil {
		return nil, errors.New("Error running revive: " + reviveError)
	}

	result := parseOutput(reviveOutput)
	return result, nil
}
