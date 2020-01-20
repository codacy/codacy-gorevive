package main

import (
	"testing"
)

func TestMapToTOML(t *testing.T) {
	expectedString := `
[rule.bar]
  arguments = [1,2]

[rule.foo]
  arguments = ["fooz"]

[rule.foozie]
`

	patternsMap := map[string]interface{}{
		"rule.foo": map[string]interface{}{
			"arguments": []interface{}{
				"fooz",
			},
		},
		"rule.bar": map[string]interface{}{
			"arguments": []interface{}{
				1, 2,
			},
		},
		"rule.foozie": map[string]interface{}{},
	}

	tomlString, err := mapToTOML(patternsMap)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	if tomlString != expectedString {
		t.Errorf("Expected toml: %s, got %s", expectedString, tomlString)
	}
}
