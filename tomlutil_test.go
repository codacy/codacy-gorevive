package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapToTOML(t *testing.T) {
	expectedString := `
[rule]

  [rule.bar]
    arguments = [1, 2]

  [rule.foo]
    arguments = ["fooz"]

  [rule.foozie]
`

	patternsMap := map[string]interface{}{
		"rule": map[string]interface{}{
			"foo": map[string]interface{}{
				"arguments": []interface{}{
					"fooz",
				},
			},
			"bar": map[string]interface{}{
				"arguments": []interface{}{
					1, 2,
				},
			},
			"foozie": map[string]interface{}{},
		},
	}

	tomlString, err := mapToTOML(patternsMap)

	assert.Nil(t, err)
	assert.Equal(t, expectedString, tomlString)
}
