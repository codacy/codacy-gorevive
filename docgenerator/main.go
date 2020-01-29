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

var docFolder string

func getMarkdownFile() *os.File {
	mdFile, err := getRulesDescriptionMarkdown()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return mdFile
}

func main() {
	flag.StringVar(&docFolder, "docFolder", "docs", "Tool documentation folder")
	flag.Parse()

	mdFile := getMarkdownFile()
	defer os.Remove(mdFile.Name())

	htmlDocumentation := convertMarkdownIntoHTML(mdFile.Name())

	patternsList := getPatternsListFromDocumentationHTML(htmlDocumentation)

	toolDefinition := createPatternsJSONFile(patternsList)

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

func createPatternsJSONFile(patterns []codacy.Pattern) codacy.ToolDefinition {
	fmt.Println("Creating patterns.json file...")

	tool := codacy.ToolDefinition{
		Name:     "codacy-gorevive",
		Version:  "1.0.0",
		Patterns: patterns,
	}

	toolAsJSON, _ := tool.ToJSONBeautify()

	ioutil.WriteFile(path.Join(docFolder, "patterns.json"), toolAsJSON, 0644)

	return tool
}

func getRuleInformationRegex(ruleName string) (*regexp.Regexp, error) {
	return regexp.Compile("(## " + ruleName + ")([\\n\\S\\s\\w\\d\\_][^##])*")
}

func getRuleDescription(ruleMdInfo string) string {
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

		ruleInformationMd := ruleInformationRegex.FindAllString(string(markdownContent), -1)

		patternDescription := codacy.PatternDescription{
			PatternID:   pattern.PatternID,
			Description: getRuleDescription(ruleInformationMd[0]),
			Title:       pattern.PatternID,
			Parameters:  pattern.Parameters,
		}

		patternsDescriptionsList = append(
			patternsDescriptionsList,
			patternDescription,
		)
		ioutil.WriteFile(
			path.Join(
				docFolder,
				"description",
				pattern.PatternID+".md",
			),
			[]byte(ruleInformationMd[0]),
			0644,
		)
	}

	descriptionsJSON, _ := json.MarshalIndent(patternsDescriptionsList, "", "  ")

	ioutil.WriteFile(
		path.Join(docFolder, "description", "description.json"),
		descriptionsJSON,
		0644,
	)
}
