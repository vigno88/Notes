package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var context, branch, date string
var description []string
var urlList []Url
var appConfigEnv []string
var hostName []string
var screenDefs [][]string
var screenDef []string

var regScreenDef = regexp.MustCompile(`ScreenDef ID:\s.*`)

// There is two html report to parse, one is hosted on a http server and the other is a
// straightup file on a server.
// To parse them, we first get them from their respective places, then parse them into Html Nodes.
// This big tree of nodes is then loopthrough and we get the different node we want.

// This function parse and get the html file from an http server
func parseItaFlowReport(url string) (urlList []Url) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("Error qetting the url" + url)
	}
	urlListNode := htmlquery.Find(doc, `//tr[@class="testRow statusFAILED"]/td[@class="dataCol linkCol dataCol7"]/a`)
	for _ , doc = range urlListNode {
		for _, attr := range doc.Attr {
			if attr.Key == "href" {
				url := Url{ attr.Val}
				urlList = append(urlList, url)
			}
		}
	}

	return urlList
}

// This function parse and get the html file from a server
func parseItaExecutionReport() {
	for _, url := range urlList {
		content, err := ioutil.ReadFile(url.UrlOfReport)
		if err != nil {
			fmt.Println("Error reading the file")
		}
		doc, err := html.Parse(strings.NewReader(string(content)))
		if err != nil {
			fmt.Println("Error parsing html")
		}

		loopItaExecutionReport(doc)
		screenDefs = append(screenDefs, screenDef)
		screenDef = nil
	}
}

// This function loop through the html node tree recursively and get the nodes we are looking for
func loopItaFlowReport(node *html.Node) {
	findURLDescription(node)
	findContextBranchDate(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		loopItaFlowReport(c)
	}
}

// This function loop through the html node tree recursively and get the nodes we are looking for
func loopItaExecutionReport(node *html.Node) {
	findTableOfAppConfigEnvHostName(node)
	findLiTagScreenDef(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		loopItaExecutionReport(c)
	}
}

// This function find the li Tag of the failed keyword
func findLiTagScreenDef(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "li" {
		for _, attr := range node.Attr {
			if attr.Val == "fail" || attr.Val == "fatal" {
				loopLi(node)
			}
		}
	}
}

// This function loops through the li tag node tree and get the node we are looking for
func loopLi(node *html.Node) {
	findScreenDef(node)
	for d := node.FirstChild; d != nil; d = d.NextSibling {
		loopLi(d)
	}
}

func findScreenDef(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "td" {
		for _, attr := range node.Attr {
			if attr.Val == "step-details" {
				if regScreenDef.MatchString(renderNode(node)) {
					content := renderNode(node)
					content = strings.TrimPrefix(content, "ScreenDef ID: ")
					content = strings.TrimSuffix(content, ".")
					screenDef = append(screenDef, content)
				}
			}
		}
	}
}

// This function get the table containing the AppConfigEnv and hostmane info
func findTableOfAppConfigEnvHostName(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "div" {
		for _, attr := range node.Attr {
			if attr.Val == "card-panel dashboard-environment" {
				table := node.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling
				loopTableDashboard(table)
			}
		}
	}
}

// This function loops through the table node tree and get the node we are looking for
func loopTableDashboard(node *html.Node) {
	findAppConfigEnvHostName(node)
	for d := node.FirstChild; d != nil; d = d.NextSibling {
		loopTableDashboard(d)
	}
}

// This function get the AppConfigEnv and Hostname nodes while looping through the tree
func findAppConfigEnvHostName(node *html.Node) {
	if node.Data == "td" {
		if renderNode(node) == "AppConfigEnv" {
			appConfigEnv = append(appConfigEnv, renderNode(node.NextSibling.NextSibling))
		}
		if renderNode(node) == "Hostname" {
			hostName = append(hostName, (renderNode(node.NextSibling.NextSibling)))
		}
	}
}

// This function get the Url and Description nodes while looping through the tree
func findURLDescription(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		for _, attr := range node.Attr {
			if attr.Val == "testRow statusFAILED" {
				for d := node.FirstChild; d != nil; d = d.NextSibling {
					for _, attr := range d.Attr {
						if attr.Val == "dataCol linkCol dataCol7" {
							a := d.FirstChild
							for _, attr := range a.Attr {
								if attr.Key == "href" {
									url := Url{}
									url.UrlOfReport = attr.Val
									urlList = append(urlList, url)
								}
							}
						}
						if attr.Val == "dataCol textCol dataCol2" {
							description = append(description, renderNode(d))
						}

					}
				}
			}
		}
	}
}

// This function get the Context, Branch and date nodes while looping through the tree
func findContextBranchDate(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "table" {
		for _, attr := range node.Attr {
			if attr.Val == "flowDataTable" {
				tBody := node.LastChild
				tr1 := tBody.FirstChild
				tr2 := tr1.NextSibling.NextSibling
				tdBranch := tr2.FirstChild.NextSibling.NextSibling.NextSibling
				branch = renderNode(tdBranch)
				tdContext := tdBranch.NextSibling.NextSibling
				context = renderNode(tdContext)
				tdDate := tdContext.NextSibling.NextSibling
				date = renderNode(tdDate)
			}
		}
	}
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return removeTag(buf.String())
}

func removeTag(node string) string {
	openingTag := regexp.MustCompile(`<\s*[^>]*>`)
	node = openingTag.ReplaceAllString(node, "")
	return node
}

func fillReports() (list []Report) {
	list = []Report{}
	for i := 0; i < len(description); i++ {
		report := Report{}
		report.TestSuite = context
		report.TestFailing = description[i]
		report.Branch = branch
		report.Date = date
		report.Environment = appConfigEnv[i]
		report.ReportUrl = urlList[i].UrlOfReport
		report.ExeHost = hostName[i]
		report.ScreenDef = screenDefs[i]
		list = append(list, report)
	}
	return list
}

func listURLToCSV(list []Url) {

	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Cannot create file")
		file.Close()
	}

	writer := csv.NewWriter(file)

	for _, url := range list {
		row := []string{url.UrlOfReport}
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Cannot write to file")
		}
	}
	writer.Flush()
	file.Close()

}
