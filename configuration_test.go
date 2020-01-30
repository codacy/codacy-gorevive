package main

import (
	"fmt"
	codacy "github.com/codacy/codacy-engine-golang-seed"
	"github.com/stretchr/testify/assert"
	"testing"
)

func patternStruct() (codacy.Pattern, string) {
	return codacy.Pattern{
			PatternID: "testing",
			Parameters: []codacy.PatternParameter{
				codacy.PatternParameter{
					Name:  "param1",
					Value: "value1",
				}, codacy.PatternParameter{
					Name:  "param2",
					Value: "value2",
				}},
			Category: "UnusedCode",
			Level:    "Info",
		}, `
[rule.testing]

  [[rule.testing.arguments]]
    param1 = "value1"
    param2 = "value2"
`
}

func TestPatternsToToml(t *testing.T) {
	pattern, _ := patternStruct()

	reviveConfigMap := patternsToReviveConfigMap([]codacy.Pattern{pattern})
	assert.NotNil(t, reviveConfigMap[reviveRuleName(pattern.PatternID)], "Expected toml: %s to exist", reviveRuleName(pattern.PatternID))
}

func TestConfigurationContentGeneration(t *testing.T) {
	pattern, patternTomlExpected := patternStruct()
	patterns := []codacy.Pattern{
		pattern,
		pattern,
	}

	configContent := generateToolConfigurationContent(patterns)
	expectedContent := fmt.Sprintf("%s", patternTomlExpected)

	assert.Equal(t, expectedContent, configContent)
}
