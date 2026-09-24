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

	// The rules table of contents is rendered as the first <ul> in the
	// document (right after the "<!-- toc -->" marker). Older revive
	// versions nested it one level deeper (ul > ul > li > a); newer
	// versions render it as a single flat list (ul > li > a). Selecting
	// the first top-level <ul> and grabbing all its <a> descendants
	// handles both shapes.
	doc.Find("ul").First().Find("a").Each(func(index int, rowhtml *goquery.Selection) {
		patternID := rowhtml.Text()
		// The TOC can also link to non-rule sections (e.g. "Configuration
		// options format"); real rule names are single kebab-case tokens
		// with no whitespace, so skip anything else.
		if strings.ContainsAny(patternID, " \n\t") {
			return
		}
		_, enabledByDefault := defaultPatterns[patternID]

		patterns = append(
			patterns,
			codacy.Pattern{
				ID:         patternID,
				Category:   "CodeStyle",
				Level:      "Info",
				Parameters: toolparameters.GetParametersForPattern(patternID),
				Enabled:    enabledByDefault,
			},
		)
	})

	return patterns, nil
}
