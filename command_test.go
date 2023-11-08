package main

import (
	"io/ioutil"
	"os"
	"testing"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
	"github.com/stretchr/testify/assert"
)

func TestParseOutput(t *testing.T) {
	reviveOutput := `{"Severity":"warning","Failure":"exported function Public should have comment or be unexported","RuleName":"exported","Category":"comments","Position":{"Start":{"Filename":"foo.go","Offset":39,"Line":7,"Column":1},"End":{"Filename":"foo.go","Offset":105,"Line":9,"Column":2}},"Confidence":1,"ReplacementLine":""}`

	parsedOutput := parseOutput(reviveOutput)

	expectedOutput := []codacy.Issue{
		{
			File:      "foo.go",
			Message:   "exported function Public should have comment or be unexported",
			PatternID: "exported",
			Line:      7,
		},
	}

	assert.Equal(t, expectedOutput, parsedOutput)
}

func TestReviveCommand(t *testing.T) {
	filesToAnalyse := []string{"foo.go", "bar.go"}
	sourceDir := "/src"
	configFile, err := ioutil.TempFile(os.TempDir(), "gorevive-")
	assert.Nil(t, err)

	cmd := reviveCommand(configFile, filesToAnalyse, sourceDir)

	assert.Equal(t,
		cmd.Args,
		[]string{"revive", "-formatter", "ndjson", "-config", configFile.Name(), "foo.go", "bar.go"},
	)
}
