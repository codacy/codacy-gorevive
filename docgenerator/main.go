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

func getMarkdownFile(toolVersion string) *os.File {
	mdFile, err := getRulesDescriptionMarkdown(toolVersion)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return mdFile
}

func readToolVersion() string {
	versionBytes, err := ioutil.ReadFile(toolVersionFile)
	if err != nil {
		return "0.0.0"
	}

	return strings.Trim(string(versionBytes), "\n")
}

func main() {
	flag.StringVar(&docFolder, "docFolder", "docs", "Tool documentation folder")
	flag.Parse()

	toolVersion := readToolVersion()

	mdFile := getMarkdownFile(toolVersion)
	defer os.Remove(mdFile.Name())

	htmlDocumentation := convertMarkdownIntoHTML(mdFile.Name())

	patternsList := getPatternsListFromDocumentationHTML(htmlDocumentation)

	toolDefinition := createPatternsJSONFile(patternsList, toolVersion)

	createDescriptionFiles(mdFile, toolDefinition.Patterns)
}

func convertMarkdownIntoHTML(markdownFileLocation string) string {
	cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html", markdownFileLocation)
	htmlConversion, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return "<body>" + string(htmlConversion) + "<body>"
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

func readMarkdownContent(mdFile *os.File) string {
	markdownContent, err := ioutil.ReadFile(mdFile.Name())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(markdownContent)
}

func createDescriptionFiles(mdFile *os.File, rulesList []codacy.Pattern) {
	fmt.Println("Creating description files...")

	markdownContent := readMarkdownContent(mdFile)

	var patternsDescriptionsList []codacy.PatternDescription

	for _, pattern := range rulesList {
		ruleInformationRegex, err := getRuleInformationRegex(pattern.PatternID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
}
