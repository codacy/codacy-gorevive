package toolparameters

import (
	"fmt"
	"strconv"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
)

// Enum for RuleParameter Type
type RuleParameterType string

const (
	//SliceType  RuleParameterType = "slice"
	ListType   RuleParameterType = "list"
	IntType    RuleParameterType = "int"
	FloatType  RuleParameterType = "float"
	StringType RuleParameterType = "string"
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
	Type RuleParameterType
}

const (
	unnamedParamName = "unnamedParam"
)

var ruleParameters = []RuleParameter{
	{
		Name: "add-constant",
		Parameters: []RuleParameter{
			{
				Name:        "maxLitCount",
				Description: "(string) maximum number of instances of a string literal that are tolerated before warn",
				Default:     "4",
				Type:        StringType,
			}, {
				Name:        "allowFloats",
				Description: "(string) comma-separated list of allowed floats",
				Default:     "",
				Type:        StringType,
			}, {
				Name:        "allowInts",
				Description: "(string) comma-separated list of allowed integers",
				Default:     "",
				Type:        StringType,
			}, {
				Name:        "allowStrs",
				Description: "(string) comma-separated list of allowed string literals",
				Default:     "",
				Type:        StringType,
			}, {
				Name:        "ignoreFuncs",
				Description: "(string) comma-separated list of function names regexp patterns to exclude",
				Default:     "os\\\\.*,fmt\\\\.Println,make",
				Type:        StringType,
			},
		},
	}, {
		Name:        "argument-limit",
		Description: "(int) the maximum number of parameters allowed per function",
		Default:     4,
		Type:        IntType,
	}, {
		Name:        "banned-characters",
		Description: "comma-separated list of character to ban",
		Default:     "",
		Type:        StringType,
	}, {
		Name:        "cognitive-complexity",
		Description: "(int) the maximum function complexity",
		Default:     7,
		Type:        IntType,
	}, {
		Name:        "comment-spacings",
		Description: "(list of string) comma-separated list of exceptions",
		Default:     "",
		Type:        ListType,
	}, {
		Name: "context-as-argument",
		Parameters: []RuleParameter{
			{
				Name:        "allowTypesBefore",
				Description: "(list of string) comma-separated list of exceptions",
				Default:     "",
				Type:        ListType,
			},
		},
	}, {
		Name:        "cyclomatic",
		Description: "(int) the maximum function complexity",
		Default:     3,
		Type:        IntType,
	}, {
		Name:        "defer",
		Description: "(string) the unused params in functions",
		Default:     "\"call-chain\",\"loop\"",
		Type:        ListType,
	}, {
		Name:        "file-header",
		Description: "(string) the header to look for in source files",
		Default:     "",
		Type:        StringType,
	}, {
		Name: "function-lenght",
		Parameters: []RuleParameter{
			{
				Name:        "maxStmt",
				Description: "(int) maximum number of statements per function",
				Default:     50,
				Type:        IntType,
			}, {
				Name:        "maxLines",
				Description: "(int) maximum number of lines per function",
				Default:     75,
				Type:        IntType,
			},
		},
	}, {
		Name:        "function-result-limit",
		Description: "(int) the maximum allowed return values",
		Default:     3,
		Type:        IntType,
	}, {
		Name:        "imports-blacklist",
		Description: "(list of string) black-list of package names (can be represented using regexp)",
		Default:     "\"crypto/md5\", \"crypto/sha1\"",
		Type:        ListType,
	}, {
		Name:        "line-length-limit",
		Description: "(int) maximum line length in characters",
		Default:     80,
		Type:        IntType,
	}, {
		Name:        "max-public-structs",
		Description: "(int) the maximum allowed public structs",
		Default:     3,
		Type:        IntType,
	}, {
		Name:        "superfluous-else",
		Description: "(list of string) the flags (\"preserveScope\")",
		Default:     "",
		Type:        ListType,
	}, {
		Name:        "unhandled-error",
		Description: "(list of string) function names regexp patterns to ignore",
		Default:     "\"fmt.Printf\"",
		Type:        ListType,
	},

	// , RuleParameter{
	// 	Name:        "string-format",
	// 	Description: "(list of string) each set is a slice containing 2-3 strings: a scope, a regex, and an optional error message",
	// 	Default:     "",
	// 	Type:        ListType,
	// }, RuleParameter{

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
