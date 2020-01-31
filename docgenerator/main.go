package main

import (
	"encoding/json"
	"flag"
	"fmt"
	codacy "github.com/codacy/codacy-engine-golang-seed"

	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const (
	toolVersionFile = ".version"
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
	versionBytes, err := ioutil.ReadFile(toolVersionFile)
	if err != nil {
		return "0.0.0"
	}

	return strings.Trim(string(versionBytes), "\n")
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

	patternsList, err := getPatternsListFromDocumentationHTML(htmlDocumentation)
	if err != nil {
		fmt.Println(err.Error())
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
		Name:     "codacy-gorevive",
		Version:  toolVersion,
		Patterns: patterns,
	}

	toolAsJSON, _ := tool.ToJSONBeautify()

	ioutil.WriteFile(path.Join(docFolder, "patterns.json"), toolAsJSON, 0644)

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
	markdownContent, err := ioutil.ReadFile(mdFile.Name())
	if err != nil {
		return "", err
	}
	return string(markdownContent), nil
}

func createDescriptionFiles(mdFile *os.File, rulesList []codacy.Pattern) error {
	fmt.Println("Creating description files...")

	markdownContent, err := readMarkdownContent(mdFile)
	if err != nil {
		return err
	}

	var patternsDescriptionsList []codacy.PatternDescription

	for _, pattern := range rulesList {
		ruleInformationRegex, err := getRuleInformationRegex(pattern.PatternID)
		if err != nil {
			return err
		}

		ruleInformationMdList := ruleInformationRegex.FindAllString(string(markdownContent), -1)
		ruleInformationMd := ""
		if len(ruleInformationMdList) > 0 {
			ruleInformationMd = strings.TrimSuffix(ruleInformationMdList[0], "##")
		}

		patternDescription := codacy.PatternDescription{
			PatternID:   pattern.PatternID,
			Description: getRuleDescription(ruleInformationMd),
			Title:       pattern.PatternID,
			Parameters:  pattern.Parameters,
		}

		patternsDescriptionsList = append(
			patternsDescriptionsList,
			patternDescription,
		)
		if len(ruleInformationMd) > 0 {
			ioutil.WriteFile(
				path.Join(
					docFolder,
					"description",
					pattern.PatternID+".md",
				),
				[]byte(ruleInformationMd),
				0644,
			)
		}
	}

	descriptionsJSON, _ := json.MarshalIndent(patternsDescriptionsList, "", "  ")

	ioutil.WriteFile(
		path.Join(docFolder, "description", "description.json"),
		descriptionsJSON,
		0644,
	)

	return nil
}
