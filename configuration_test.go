package main

import (
	"fmt"
	codacy "github.com/josemiguelmelo/codacy-engine-golang-seed"
	"testing"
)

func patternStruct() (codacy.Pattern, string) {
	return codacy.Pattern{
			PatternID: "testing",
			Parameters: []codacy.PatternParameter{
				codacy.PatternParameter{
					Name:  "param1",
					Value: "value1",
				}},
			Category: "UnusedCode",
			Level:    "Info",
		}, `[rule.testing]
  arguments = [{param1 = "value1"}]`
}

func TestPatternToToml(t *testing.T) {
	pattern, expectedToml := patternStruct()

	patternsAsToml := patternToToml(pattern)

	if patternsAsToml != expectedToml {
		t.Errorf("Expected toml: %s, got %s", expectedToml, patternsAsToml)
	}
}

func TestReviveTomlConfigurationContent(t *testing.T) {
	pattern, patternTomlExpected := patternStruct()
	patterns := []codacy.Pattern{
		pattern,
		pattern,
	}

	configContent := reviveTomlConfigurationContent(patterns)
	expectedContent := fmt.Sprintf("%s\n\n%s\n\n", patternTomlExpected, patternTomlExpected)

	if configContent != expectedContent {
		t.Errorf("Expected toml configuration: %s, got %s", expectedContent, configContent)
	}
}
