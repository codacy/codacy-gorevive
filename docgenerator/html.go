package main

import (
	"strings"

	toolparameters "codacy.com/codacy-gorevive/toolparameters"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	codacy "github.com/codacy/codacy-golang-tools-engine"
	"os"
)

func getPatternsListFromDocumentationHTML(data string) []codacy.Pattern {
	patterns := []codacy.Pattern{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	doc.Find("ul").Each(func(index int, externalUlHtml *goquery.Selection) {
		externalUlHtml.Find("ul").Each(func(index int, internalUlHtml *goquery.Selection) {
			internalUlHtml.Find("li").Each(func(index int, tablehtml *goquery.Selection) {
				tablehtml.Find("a").Each(func(indextr int, rowhtml *goquery.Selection) {
					patternID := rowhtml.Text()
					patterns = append(
						patterns,
						codacy.Pattern{
							PatternID:  patternID,
							Category:   "CodeStyle",
							Level:      "Info",
							Parameters: toolparameters.GetParametersForPattern(patternID),
						},
					)
				})
			})
		})
	})

	return patterns
}
