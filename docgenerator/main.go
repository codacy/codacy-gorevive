package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"

	"strings"

	"github.com/PuerkitoBio/goquery"
	codacy "github.com/josemiguelmelo/codacy-engine-golang-seed"
)

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string) (*os.File, error) {
	// Create the file
	out, err := ioutil.TempFile(os.TempDir(), "tmp-gorevive-")
	if err != nil {
		return nil, err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getPatternsFromDocumentationHtml(data string) []codacy.Pattern {
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
					patterns = append(patterns,
						codacy.Pattern{
							PatternID: rowhtml.Text(),
							Category:  "CodeStyle",
							Level:     "Info",
						})
				})
			})
		})
	})
	return patterns
}

func main() {
	docFolder := flag.String("docFolder", "docs", "Tool documentation folder")
	flag.Parse()

	mdFile, err := DownloadFile("https://raw.githubusercontent.com/mgechev/revive/master/RULES_DESCRIPTIONS.md")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer os.Remove(mdFile.Name())

	cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html", mdFile.Name())
	htmlConversion, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	htmlDocumentation := "<body>" + string(htmlConversion) + "<body>"

	tool := codacy.ToolDefinition{
		Name:     "codacy-gorevive",
		Version:  "1.0.0",
		Patterns: getPatternsFromDocumentationHtml(htmlDocumentation),
	}

	toolAsJSON, _ := tool.ToJSONBeautify()

	ioutil.WriteFile(path.Join(*docFolder, "patterns.json"), toolAsJSON, 0644)
}
