package main

import (
	"fmt"
	"testing"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
	"github.com/stretchr/testify/assert"
)

func patternStruct() (codacy.Pattern, string) {
	return codacy.Pattern{
			ID: "testing",
			Parameters: []codacy.PatternParameter{
				{
					Name:  "param1",
					Value: "value1",
				},
				{
					Name:  "param2",
					Value: "value2",
				}},
			Category: "UnusedCode",
			Level:    "Info",
		}, `
[rule]

  [rule.testing]

    [[rule.testing.arguments]]
      param1 = "value1"
      param2 = "value2"
`
}

func patternUnnamedStruct() (codacy.Pattern, string) {
	return codacy.Pattern{
			ID: "testingUnnamed",
			Parameters: []codacy.PatternParameter{
				{
					Name:  unnamedParamName,
					Value: "value1",
				},
				{
					Name:  "param2",
					Value: "value2",
				}},
			Category: "UnusedCode",
			Level:    "Info",
		}, `
  [rule.testingUnnamed]
    arguments = ["value1"]
`
}

func TestPatternsToToml(t *testing.T) {
	pattern, _ := patternStruct()

	reviveConfigMap := patternsToReviveConfigMap([]codacy.Pattern{pattern})
	value, ok := reviveConfigMap["rule"].(map[string]interface{})
	if !ok {
		assert.Fail(t, "Revive config map does not have the expected structure")
	}
	assert.NotNil(t, value[pattern.ID], "Expected toml: %s to exist", pattern.ID)
}

func TestConfigurationContentGeneration(t *testing.T) {
	pattern, patternTomlExpected := patternStruct()
	unnamedParamPattern, unnamedParamTomlExpected := patternUnnamedStruct()
	patterns := []codacy.Pattern{
		pattern,
		unnamedParamPattern,
	}

	configContent := generateToolConfigurationContent(patterns)

	expectedContent := fmt.Sprintf("%s%s", patternTomlExpected, unnamedParamTomlExpected)

	assert.Equal(t, expectedContent, configContent)
}
