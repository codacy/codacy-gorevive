package main

import (
	"encoding/json"
	"os/exec"

	codacy "github.com/codacy/codacy-engine-golang-seed"

	"os"
	"strings"
)

const (
	formatter = "ndjson"
)

// ReviveResult base structure of revive json result
type ReviveResult struct {
	Failure  string
	RuleName string
	Position RevivePosition
}

// RevivePosition revive json position
type RevivePosition struct {
	Start RevivePositionResult
	End   RevivePositionResult
}

// RevivePositionResult revive json position result
type RevivePositionResult struct {
	Filename string
	Line     int
	Offset   int
	Column   int
}

func getConfigFileParam(configFile *os.File) []string {
	if configFile != nil {
		return []string{
			"-config",
			configFile.Name(),
		}
	}
	return []string{}
}

func commandParameters(configFile *os.File, filesToAnalyse []string) []string {
	cmdParams := append(
		[]string{
			"-formatter", formatter,
		},
		getConfigFileParam(configFile)...,
	)

	cmdParams = append(cmdParams, filesToAnalyse...)

	return cmdParams
}

func parseOutput(reviveOutput string) []codacy.Issue {
	var result []codacy.Issue
	lines := strings.Split(reviveOutput, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		var reviveRes ReviveResult
		json.Unmarshal([]byte(line), &reviveRes)

		result = append(result, codacy.Issue{
			PatternID: reviveRes.RuleName,
			Message:   reviveRes.Failure,
			Line:      reviveRes.Position.Start.Line,
			File:      reviveRes.Position.Start.Filename,
		})
	}
	return result
}

func reviveCommand(configFile *os.File, filesToAnalyse []string, sourceDir string) *exec.Cmd {
	params := commandParameters(configFile, filesToAnalyse)

	cmd := exec.Command("revive", params...)

	cmd.Dir = sourceDir
	return cmd
}
