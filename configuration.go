package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	toolparameters "codacy.com/codacy-gorevive/toolparameters"
	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
)

const (
	unnamedParamName     = "unnamedParam"
	sourceConfigFileName = "revive.toml"
)

// paramValueByType checks the type of parameter according to the tool documentation
func paramValueByType(paramValue interface{}, ruleDefinition toolparameters.RuleParameter) interface{} {

	switch ruleDefinition.Type {
	case toolparameters.ListType:
		return strings.Split(paramValue.(string), ", ")
	case toolparameters.IntType:
		return int(paramValue.(float64))
	case toolparameters.FloatType:
		return paramValue.(float64)
	case toolparameters.StringType:
		return fmt.Sprintf("%v", paramValue)
	default:
		return paramValue
	}
}

// paramValue converts codacy's parameter into a revive parameter value
func paramValue(param codacy.PatternParameter, patternID string) interface{} {
	ruleDefinition, err := toolparameters.FindRuleParameterDefinition(patternID)
	if err != nil {
		// fallback: rule not found
		switch v := param.Value.(type) {
		case float64:
			return int(v)
		case string:
			if i, convErr := strconv.Atoi(v); convErr == nil {
				return i
			}
			return v
		default:
			return param.Value
		}
	}

	// normal case: ruleDefinition was found
	for _, p := range ruleDefinition.Parameters {
		if p.Name == param.Name {
			if param.Value != nil {
				return paramValueByType(param.Value, p)
			}
			return paramValueByType(p.Default, p)
		}
	}

	if param.Value == nil {
		return paramValueByType(param.Default, ruleDefinition)
	}
	return paramValueByType(param.Value, ruleDefinition)
}

func unnamedParam(value interface{}) []interface{} {
	resultTmp := []interface{}{}
	switch value.(type) {
	case []string:
		// if is a []string, append all values to res, one by one
		for _, v := range value.([]string) {
			resultTmp = append(resultTmp, v)
		}
	default:
		resultTmp = append(resultTmp, value)
	}
	return resultTmp
}

// patternParametersAsReviveValues converts pattern parameters into a list of revive arguments
func patternParametersAsReviveValues(pattern codacy.Pattern) []interface{} {
	namedParameters := map[string]interface{}{}
	for _, p := range pattern.Parameters {
		value := paramValue(p, pattern.ID)

		if p.Name == unnamedParamName {
			return unnamedParam(value)
		}
		namedParameters[p.Name] = value
	}

	if len(namedParameters) > 0 {
		return []interface{}{
			namedParameters,
		}
	}

	return []interface{}{}
}

func reviveArguments(paramsValues []interface{}) map[string]interface{} {
	if paramsValues == nil || len(paramsValues) == 0 {
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"arguments": paramsValues,
	}
}

func patternsToReviveConfigMap(patterns []codacy.Pattern) map[string]interface{} {
	patternsMap := map[string]interface{}{}
	for _, pattern := range patterns {
		patternsMap[pattern.ID] = reviveArguments(patternParametersAsReviveValues(pattern))
	}

	rules := map[string]interface{}{
		"rule": patternsMap,
	}
	return rules
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

func configurationFromSourceCode(sourceFolder string) (string, error) {
	filename := path.Join(sourceFolder, sourceConfigFileName)
	contentByte, err := os.ReadFile(filename)
	return string(contentByte), err
}

// getConfigurationFile returns file, boolean saying if it is temp and error
func getConfigurationFile(patterns *[]codacy.Pattern, sourceFolder string) (*os.File, error) {
	// if no patterns, try to use configuration from source code
	// otherwise default configuration file
	if patterns == nil {
		sourceConfigFileContent, err := configurationFromSourceCode(sourceFolder)
		if err == nil {
			return writeToTempFile(sourceConfigFileContent)
		}

		return nil, err
	}

	content := generateToolConfigurationContent(*patterns)

	return writeToTempFile(content)
}
