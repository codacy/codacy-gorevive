package main

import (
	"strings"

	toolparameters "codacy.com/codacy-gorevive/toolparameters"
	"github.com/PuerkitoBio/goquery"
	codacy "github.com/codacy/codacy-engine-golang-seed/v6"
)

func getPatternsListFromDocumentationHTML(data string, defaultPatterns map[string]interface{}) ([]codacy.Pattern, error) {
	patterns := []codacy.Pattern{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	doc.Find("ul").Each(func(index int, externalUlHtml *goquery.Selection) {
		externalUlHtml.Find("ul").Each(func(index int, internalUlHtml *goquery.Selection) {
			internalUlHtml.Find("li").Each(func(index int, tablehtml *goquery.Selection) {
				tablehtml.Find("a").Each(func(indextr int, rowhtml *goquery.Selection) {
					patternID := rowhtml.Text()
					_, enabledByDefault := defaultPatterns[patternID]

					patterns = append(
						patterns,
						codacy.Pattern{
							PatternID:  patternID,
							Category:   "CodeStyle",
							Level:      "Info",
							Parameters: toolparameters.GetParametersForPattern(patternID),
							Enabled:    enabledByDefault,
						},
					)
				})
			})
		})
	})

	return patterns, nil
}
