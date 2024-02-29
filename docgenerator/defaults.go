package main

import (
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

func getDefaultPatterns(toolVersion string) map[string]interface{} {
	defaultPatternsFile, err := downloadDefaultsToml(toolVersion)
	if err != nil {
		fmt.Println("Error downloading defaults.toml file:", err)
		return nil
	}

	return readRulesFromToml(defaultPatternsFile)
}

func readRulesFromToml(tomlFile *os.File) map[string]interface{} {
	tomlTree, err := toml.LoadFile(tomlFile.Name())
	if err != nil {
		fmt.Println("Error reading toml file:", err)
		return nil
	}

	rulesMap := tomlTree.ToMap()["rule"].(map[string]interface{})
	return rulesMap
}
