package main

import (
	toml "github.com/pelletier/go-toml/v2"
)

func mapToTOML(jsonMap map[string]interface{}) (string, error) {
	tree, err := toml.TreeFromMap(jsonMap)
	if err != nil {
		return "", err
	}

	tomlBytes, err := tree.ToTomlString()
	if err != nil {
		return "", err
	}
	return string(tomlBytes[:]), nil
}
