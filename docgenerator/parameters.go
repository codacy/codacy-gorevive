package main

import (
	codacy "github.com/codacy/codacy-golang-tools-engine"
)

type RuleParameter struct {
	Name        string
	Parameters  map[string]string
	Description string
}

const (
	unnamedParamName = "unnamedParam"
)

var ruleParameters = []RuleParameter{
	RuleParameter{
		Name:        "file-header",
		Description: "(string) the header to look for in source files.",
	}, RuleParameter{
		Name: "add-constant",
		Parameters: map[string]string{
			"allowFloats": "(string) comma-separated list of allowed floats",
			"allowInts":   "(string) comma-separated list of allowed integers",
			"allowStrs":   "(string) comma-separated list of allowed string literals",
			"maxLitCount": "(string) maximum number of instances of a string literal that are tolerated before warn.",
		},
	}, RuleParameter{
		Name:        "imports-blacklist",
		Description: "(list of string) black-list of package names",
	}, RuleParameter{
		Name:        "line-length-limit",
		Description: "(int) maximum line length in characters.",
	}, RuleParameter{
		Name:        "function-result-limit",
		Description: "(int) the maximum allowed return values",
	}, RuleParameter{
		Name:        "max-public-structs",
		Description: "(int) the maximum allowed public structs",
	}, RuleParameter{
		Name:        "cyclomatic",
		Description: "(int) the maximum function complexity",
	}, RuleParameter{
		Name:        "argument-limit",
		Description: "(int) the maximum number of parameters allowed per function.",
	}, RuleParameter{
		Name:        "cognitive-complexity",
		Description: "(int) the maximum function complexity",
	}, RuleParameter{
		Name:        "unhandled-error",
		Description: "(list of string) function names to ignore",
	},
}

func getParametersForPattern(patternID string) []codacy.PatternParameter {
	for _, rule := range ruleParameters {
		if rule.Name == patternID {
			var params []codacy.PatternParameter
			if rule.Parameters == nil {
				return []codacy.PatternParameter{
					codacy.PatternParameter{
						Name:        unnamedParamName,
						Description: rule.Description,
					},
				}
			}

			for paramKey, paramDescr := range rule.Parameters {
				params = append(params,
					codacy.PatternParameter{
						Name:        paramKey,
						Description: paramDescr,
					})
			}
			return params
		}
	}

	return nil
}
