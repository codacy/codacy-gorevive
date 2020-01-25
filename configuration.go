package main

import (
	"fmt"
	codacy "github.com/codacy/codacy-golang-tools-engine"
	"io/ioutil"
	"os"
	"path"
)

const (
	unnamedParamName     = "unnamedParam"
	sourceConfigFileName = "codacyrc.toml"
)

func paramValue(p interface{}) interface{} {
	if isInteger(p) {
		return int(p.(float64))
	} else {
		return p
	}
}

func patternParametersAsListOfValues(parameters []codacy.PatternParameter) []interface{} {
	var res []interface{}
	if len(parameters) == 0 {
		return []interface{}{}
	}

	namedParameters := map[string]interface{}{}
	for _, p := range parameters {
		value := paramValue(p.Value)

		if p.Name == unnamedParamName {
			res = append(res, value)
		} else {
			namedParameters[p.Name] = value
		}
	}

	if len(namedParameters) > 0 {
		res = append(res, namedParameters)
	}

	return res
}

func reviveArguments(pattern codacy.Pattern) map[string]interface{} {
	paramsValues := patternParametersAsListOfValues(pattern.Parameters)
	if paramsValues == nil || len(paramsValues) == 0 {
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"arguments": paramsValues,
	}
}

func patternsToReviveConfigMap(patterns []codacy.Pattern) map[string]interface{} {
	var patternsMap = make(map[string]interface{})
	for _, pattern := range patterns {
		patternsMap["rule."+pattern.PatternID] = reviveArguments(pattern)
	}
	return patternsMap
}

func generateToolConfigurationContent(patterns []codacy.Pattern) string {
	patternsMap := patternsToReviveConfigMap(patterns)

	tomlString, err := mapToTOML(patternsMap)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return tomlString
}

func writeConfigurationToTempFile(content string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "gorevive-")
	if err != nil {
		return nil, err
	}
	if _, err = tmpFile.Write([]byte(content)); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func getConfigurationFromSourceCode(sourceFolder string) (string, error) {
	filename := path.Join(sourceFolder, sourceConfigFileName)

	contentByte, err := ioutil.ReadFile(filename)
	return string(contentByte), err
}

// getConfigurationFile returns file, boolean saying if it is temp and error
func getConfigurationFile(patterns []codacy.Pattern, sourceFolder string) (*os.File, error) {
	// if no patterns, try to use configuration from source code
	// otherwise default configuration file
	if len(patterns) == 0 {
		sourceConfigFileContent, err := getConfigurationFromSourceCode(sourceFolder)
		if err == nil {
			return writeConfigurationToTempFile(sourceConfigFileContent)
		}

		return nil, nil
	}

	content := generateToolConfigurationContent(patterns)

	return writeConfigurationToTempFile(content)
}
