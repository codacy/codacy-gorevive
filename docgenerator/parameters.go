package main

import (
	"errors"
	"fmt"
	codacy "github.com/codacy/codacy-golang-tools-engine"
)

type RuleParameter struct {
	Name        string
	Parameters  []codacy.PatternParameter
	Default     string
	Description string
	Type        string
}

const (
	unnamedParamName = "unnamedParam"

	ListType   = "list"
	IntType    = "int"
	StringType = "string"
)

var ruleParameters = []RuleParameter{
	RuleParameter{
		Name:        "file-header",
		Description: "(string) the header to look for in source files.",
	}, RuleParameter{
		Name: "add-constant",
		Parameters: []codacy.PatternParameter{
			codacy.PatternParameter{
				Name:        "allowFloats",
				Default:     "",
				Description: "(string) comma-separated list of allowed floats",
			},
			codacy.PatternParameter{
				Name:        "allowInts",
				Default:     "",
				Description: "allowInts",
			},
			codacy.PatternParameter{
				Name:        "allowStrs",
				Default:     "",
				Description: "(string) comma-separated list of allowed string literals",
			},
			codacy.PatternParameter{
				Name:        "maxLitCount",
				Default:     "",
				Description: "(string) maximum number of instances of a string literal that are tolerated before warn.",
			},
		},
	}, RuleParameter{
		Name:        "line-length-limit",
		Description: "(int) maximum line length in characters.",
		Default:     "80",
		Type:        IntType,
	}, RuleParameter{
		Name:        "function-result-limit",
		Description: "(int) the maximum allowed return values",
		Default:     "3",
		Type:        IntType,
	}, RuleParameter{
		Name:        "max-public-structs",
		Description: "(int) the maximum allowed public structs",
		Default:     "3",
		Type:        IntType,
	}, RuleParameter{
		Name:        "cyclomatic",
		Description: "(int) the maximum function complexity",
		Default:     "3",
		Type:        IntType,
	}, RuleParameter{
		Name:        "argument-limit",
		Description: "(int) the maximum number of parameters allowed per function.",
		Default:     "4",
		Type:        IntType,
	}, RuleParameter{
		Name:        "cognitive-complexity",
		Description: "(int) the maximum function complexity",
		Default:     "7",
		Type:        IntType,
	},
	RuleParameter{
		Name:        "unhandled-error",
		Description: "(list of string) function names to ignore",
		Default:     "[\"fmt.Printf\"]",
		Type:        ListType,
	}, RuleParameter{
		Name:        "imports-blacklist",
		Description: "(list of string) black-list of package names",
		Default:     "[\"crypto/md5\", \"crypto/sha1\"]",
		Type:        ListType,
	},
}

func FindRuleParameterDefinition(patternID string) (RuleParameter, error) {
	for _, rule := range ruleParameters {
		if rule.Name == patternID {
			return rule, nil
		}
	}
	return RuleParameter{},
		errors.New(fmt.Sprintf("Not found parameters for pattern with id %s", patternID))
}

// GetParametersForPattern returns the parameters for the pattern with id patternID
func GetParametersForPattern(patternID string) []codacy.PatternParameter {
	rule, err := FindRuleParameterDefinition(patternID)
	if err != nil {
		return nil
	}

	if len(rule.Parameters) == 0 {
		return []codacy.PatternParameter{
			codacy.PatternParameter{
				Name:        unnamedParamName,
				Description: rule.Description,
				Default:     rule.Default,
			},
		}
	}

	return rule.Parameters
}
