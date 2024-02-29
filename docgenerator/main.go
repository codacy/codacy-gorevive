// generate documentation for the tool
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	codacy "github.com/codacy/codacy-engine-golang-seed/v6"

	"os"
	"os/exec"
	"regexp"
	"strings"

	stripmd "github.com/writeas/go-strip-markdown"
)

const (
	toolVersionFile = ".tool_version"
)

var docFolder string

func getMarkdownFile(toolVersion string) (*os.File, error) {
	mdFile, err := getRulesDescriptionMarkdown(toolVersion)
	if err != nil {
		return nil, err
	}
	return mdFile, nil
}

func readToolVersion() string {
	versionBytes, err := os.ReadFile(toolVersionFile)
	if err != nil {
		return "0.0.0"
	}

	return strings.Trim(string(versionBytes), "\r\n")
}

func run() int {
	toolVersion := readToolVersion()

	mdFile, err := getMarkdownFile(toolVersion)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	defer os.Remove(mdFile.Name())

	htmlDocumentation, err := convertMarkdownIntoHTML(mdFile.Name())
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	defaultPatterns := getDefaultPatterns(toolVersion)

	patternsList, err := getPatternsListFromDocumentationHTML(htmlDocumentation, defaultPatterns)
	if err != nil {
		fmt.Println("Error getting patterns list from documentation:", err)
		return 1
	}

	toolDefinition := createPatternsJSONFile(patternsList, toolVersion)

	createDescriptionFiles(mdFile, toolDefinition.Patterns)

	return 0
}

func main() {
	flag.StringVar(&docFolder, "docFolder", "docs", "Tool documentation folder")
	flag.Parse()
	os.Exit(run())
}

func convertMarkdownIntoHTML(markdownFileLocation string) (string, error) {
	cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html", markdownFileLocation)
	htmlConversion, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return "<body>" + string(htmlConversion) + "<body>", nil
}

func createPatternsJSONFile(patterns []codacy.Pattern, toolVersion string) codacy.ToolDefinition {
	fmt.Println("Creating patterns.json file...")

	tool := codacy.ToolDefinition{
		Name:     "revive",
		Version:  toolVersion,
		Patterns: patterns,
	}

	toolAsJSON, _ := json.MarshalIndent(tool, "", "  ")

	err := os.WriteFile(filepath.Join(docFolder, "patterns.json"), toolAsJSON, 0640)
	if err != nil {
		fmt.Println("Error creating patterns.json file:", err)
	}

	return tool
}

func getRuleInformationRegex(ruleName string) (*regexp.Regexp, error) {
	return regexp.Compile("(## " + ruleName + ")([\\n\\S\\s\\w\\d\\_])*?##")
}

func getRuleDescription(ruleMdInfo string) string {
	if ruleMdInfo == "" {
		return ruleMdInfo
	}
	descriptionPrefix := "_Description_: "
	descriptionRegex := regexp.MustCompile(descriptionPrefix + ".*")
	descriptionString := descriptionRegex.FindAllString(ruleMdInfo, -1)[0]
	return strings.Replace(descriptionString, descriptionPrefix, "", -1)
}

func readMarkdownContent(mdFile *os.File) (string, error) {
	markdownContent, err := os.ReadFile(mdFile.Name())
	if err != nil {
		return "", err
	}
	return string(markdownContent), nil
}

func createDescriptionFiles(mdFile *os.File, rulesList []codacy.Pattern) {
	descriptionFolder := filepath.Join(docFolder, "description")
	fmt.Println("Creating description files...")

	markdownContent, err := readMarkdownContent(mdFile)
	if err != nil {
		fmt.Println("Error reading markdown file:", err)
		return
	}

	var patternsDescriptionsList []codacy.PatternDescription

	for _, pattern := range rulesList {
		ruleInformationRegex, err := getRuleInformationRegex(pattern.ID)
		if err != nil {
			fmt.Println("Error getting rule information regex:", err)
			return
		}

		ruleInformationMdList := ruleInformationRegex.FindAllString(string(markdownContent), -1)
		ruleInformationMd := ""
		if len(ruleInformationMdList) > 0 {
			ruleInformationMd = strings.TrimSuffix(ruleInformationMdList[0], "##")
		}

		// Get description for parameters
		var params []codacy.PatternParameter
		for _, param := range pattern.Parameters {
			params = append(params, codacy.PatternParameter{
				Name:        param.Name,
				Description: param.Description,
			})
		}

		stripedDescription := stripmd.Strip(getRuleDescription(ruleInformationMd))
		// Required this replace due to bug on markdown strip for some URLs.
		urlReg, _ := regexp.Compile(`#[^)]*\)`)
		stripedDescription = string(urlReg.ReplaceAll([]byte(stripedDescription), []byte("")))

		patternDescription := codacy.PatternDescription{
			PatternID:   pattern.ID,
			Description: stripedDescription,
			Title:       pattern.ID,
			Parameters:  params,
		}

		patternsDescriptionsList = append(
			patternsDescriptionsList,
			patternDescription,
		)
		if len(ruleInformationMd) > 0 {
			path := filepath.Join(descriptionFolder, pattern.ID+".md")
			err = os.WriteFile(
				path,
				[]byte(ruleInformationMd),
				0640,
			)
			if err != nil {
				fmt.Println("Error writing pattern description file:", path, err)
			}
		}
	}

	descriptionsJSON, err := json.MarshalIndent(patternsDescriptionsList, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling descriptions to JSON:", err)
		return
	}

	path := filepath.Join(descriptionFolder, "description.json")
	err = os.WriteFile(
		path,
		descriptionsJSON,
		0640,
	)
	if err != nil {
		fmt.Println("Error writing description file:", path, err)
	}
}
