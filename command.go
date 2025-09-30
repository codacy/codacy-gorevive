package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"

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

func parseOutput(reviveOutput string) []codacy.Result {
	var result []codacy.Result

	scanner := bufio.NewScanner(strings.NewReader(reviveOutput))
	for scanner.Scan() {
		var reviveRes ReviveResult
		json.Unmarshal([]byte(scanner.Text()), &reviveRes)
		if reviveRes.RuleName == "" {
			continue
		}
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

func runCommand(cmd *exec.Cmd) (string, string, error) {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	cmdOutput, err := cmd.Output()
	if err != nil {
		return "", stderr.String(), err
	}
	return string(cmdOutput), "", nil
}
