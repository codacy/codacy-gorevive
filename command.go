package main

import (
	"encoding/json"
	"os/exec"

	codacy "github.com/codacy/codacy-golang-tools-engine"

	"os"
	"strings"
)

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

func command(configFile *os.File, filesToAnalyse []string, sourceDir string) *exec.Cmd {
	params := commandParameters(configFile, filesToAnalyse)

	cmd := exec.Command("revive", params...)

	cmd.Dir = "/src"
	return cmd
}
