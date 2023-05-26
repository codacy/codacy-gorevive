package toolparameters

import (
	"fmt"
	"strconv"

	codacy "github.com/codacy/codacy-engine-golang-seed"
)

// RuleParameter contains the definition of the parameters for a rule
type RuleParameter struct {
	// Name rule name
	Name string
	// Parameters list of parameters for the rule. If none, then the parameter will be unnamed
	Parameters []RuleParameter
	// Default value for the unnamed parameter
	Default interface{}
	// Description for the unnamed parameter
	Description string
	// Type of the unnamed parameter
	Type string
}

const (
	unnamedParamName = "unnamedParam"

	// ListType specifies the list type of param
	ListType = "list"
	// IntType specifies the int type of param
	IntType = "int"
	// FloatType specifies the float type of param
	FloatType = "float"
	// StringType specifies the string type of param
	StringType = "string"
)

var ruleParameters = []RuleParameter{
	RuleParameter{
		Name:        "file-header",
		Description: "(string) the header to look for in source files.",
		Default:     "",
		Type:        StringType,
	}, RuleParameter{
		Name: "add-constant",
		Parameters: []RuleParameter{
			{
				Name:        "allowFloats",
				Default:     "",
				Type:        StringType,
				Description: "(string) comma-separated list of allowed floats",
			},
			{
				Name:        "allowInts",
				Default:     "",
				Type:        StringType,
				Description: "allowInts",
			},
			{
				Name:        "allowStrs",
				Default:     "",
				Type:        StringType,
				Description: "(string) comma-separated list of allowed string literals",
			},
			{
				Name:        "maxLitCount",
				Default:     "4",
				Type:        StringType,
				Description: "(string) maximum number of instances of a string literal that are tolerated before warn.",
			},
		},
	}, RuleParameter{
		Name:        "line-length-limit",
		Description: "(int) maximum line length in characters.",
		Default:     80,
		Type:        IntType,
	}, RuleParameter{
		Name:        "function-result-limit",
		Description: "(int) the maximum allowed return values",
		Default:     3,
		Type:        IntType,
	}, RuleParameter{
		Name: "function-length",
		Parameters: []RuleParameter{
			{
				Name:        "maxStatements",
				Default:     10,
				Type:        IntType,
				Description: "(int) the maximum allowed statements",
			},
			{
				Name:        "maxLines",
				Default:     0,
				Type:        IntType,
				Description: "(int) the maximum allowed lines",
			},
		},
	}, RuleParameter{
		Name:        "max-public-structs",
		Description: "(int) the maximum allowed public structs",
		Default:     3,
		Type:        IntType,
	}, RuleParameter{
		Name:        "cyclomatic",
		Description: "(int) the maximum function complexity",
		Default:     3,
		Type:        IntType,
	}, RuleParameter{
		Name:        "argument-limit",
		Description: "(int) the maximum number of parameters allowed per function.",
		Default:     4,
		Type:        IntType,
	}, RuleParameter{
		Name:        "cognitive-complexity",
		Description: "(int) the maximum function complexity",
		Default:     7,
		Type:        IntType,
	},
	RuleParameter{
		Name:        "unhandled-error",
		Description: "(list of string) function names to ignore",
		Default:     "\"fmt.Printf\"",
		Type:        ListType,
	}, RuleParameter{
		Name:        "imports-blacklist",
		Description: "(list of string) black-list of package names",
		Default:     "\"crypto/md5\", \"crypto/sha1\"",
		Type:        ListType,
	},
}

// FindRuleParameterDefinition finds the parameter definition for a rule. If it does not find, an error is returned
func FindRuleParameterDefinition(patternID string) (RuleParameter, error) {
	for _, rule := range ruleParameters {
		if rule.Name == patternID {
			return rule, nil
		}
	}
	return RuleParameter{}, fmt.Errorf("Not found parameters for pattern with id %s", patternID)
}

// GetParametersForPattern returns the parameters for the pattern with id patternID
func GetParametersForPattern(patternID string) []codacy.PatternParameter {
	rule, err := FindRuleParameterDefinition(patternID)
	if err != nil {
		return nil
	}

	if len(rule.Parameters) == 0 {
		defaultValue := rule.Default
		if rule.Type == IntType {
			defaultValue = strconv.Itoa(defaultValue.(int))
		}

		return []codacy.PatternParameter{
			{
				Name:        unnamedParamName,
				Description: rule.Description,
				Default:     defaultValue,
			},
		}
	}

	parameters := []codacy.PatternParameter{}
	for _, param := range rule.Parameters {
		defaultValue := param.Default
		if rule.Type == IntType {
			defaultValue = strconv.Itoa(defaultValue.(int))
		}
		parameters = append(parameters, codacy.PatternParameter{
			Name:        param.Name,
			Description: param.Description,
			Default:     defaultValue,
		})
	}

	return parameters
}
