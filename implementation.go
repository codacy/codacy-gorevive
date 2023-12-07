package main

import (
	"errors"
	"path/filepath"
	"strings"

	"os"

	codacy "github.com/codacy/codacy-engine-golang-seed"
	"github.com/mgechev/revive/config"
	"github.com/mgechev/revive/lint"
	revivelib "github.com/mgechev/revive/revivelib"
)

// GoReviveImplementation tool implementation
type GoReviveImplementation struct {
}

// Run runs the tool implementation
func (i GoReviveImplementation) Run(tool codacy.Tool, sourceDir string) ([]codacy.Issue, error) {
	configFile, err := getConfigurationFile(tool.Patterns, sourceDir)
	if err == nil {
		defer os.Remove(configFile.Name())
	}

	filesToAnalyse, err := getListOfFilesToAnalyse(tool.Files, sourceDir)
	if err != nil {
		return nil, errors.New("Error getting files to analyse: " + err.Error())
	}

	c := ""
	if configFile != nil {
		c = configFile.Name()
	}

	configs, err := config.GetConfig(c)
	if err != nil {
		return nil, errors.New("Error parsing config file: " + err.Error())
	}

	revive, _ := revivelib.New(
		configs,
		false,
		0,
	)

	packages := []*revivelib.LintPattern{}
	for _, file := range filesToAnalyse {
		packages = append(packages, revivelib.Include(filepath.Join(sourceDir, file)))
	}

	failures, err := revive.Lint(packages...)
	if err != nil {
		return nil, errors.New("Error running revive: " + err.Error())
	}

	result := convertToCodacyAPI(sourceDir, failures)
	return result, nil
}

func convertToCodacyAPI(sourceDir string, failures <-chan lint.Failure) []codacy.Issue {
	issues := []codacy.Issue{}
	for failure := range failures {
		issue :=
			codacy.Issue{
				PatternID: failure.RuleName,
				File:      strings.Replace(failure.GetFilename(), sourceDir+"/", "", 1),
				Line:      failure.Position.Start.Line,
				Message:   failure.Failure,
			}
		issues = append(issues, issue)
	}

	return issues
}
