package main

import (
	"os"

	toml "github.com/pelletier/go-toml"
)

func getDefaultPatterns(toolVersion string) (map[string]interface{}, error) {
	defaultPatternsFile, err := downloadDefaultsToml(toolVersion)
	if err != nil {
		return nil, err
	}

	return readRulesFromToml(defaultPatternsFile)
}

func readRulesFromToml(tomlFile *os.File) (map[string]interface{}, error) {
	tomlTree, err := toml.LoadFile(tomlFile.Name())
	if err != nil {
		return nil, err
	}

	rulesMap := tomlTree.ToMap()["rule"].(map[string]interface{})
	return rulesMap, nil
}
