package main

//
//import (
//	"fmt"
//	"github.com/antchfx/htmlquery"
//	"github.com/antchfx/xmlquery"
//	"golang.org/x/net/html"
//	"strings"
//)
//func parseItaFlowReportBDI(url string) (reports []Report) {
//
//	doc, err := htmlquery.LoadURL(url)
//	if err != nil {
//		fmt.Println("Error qetting the url" + url)
//	}
//	urlList, description := getUrlAndDescription(doc)
//	date := renderNode(htmlquery.Find(doc, `//table[@class="flowDataTable"]/tbody/tr[2]/td[4]`)[0])
//
//	for i , url := range urlList {
//		var report Report
//		report.ReportUrl = url
//		report.TestFailing = description[i]
//		report.Date = date
//		report.TestSuite, report.Environment, report.ExeHost, report.ScreenDef = parseItaExecutionReportBDI(url)
//		reports = append(reports, report)
//	}
//	return reports
//}
//
//
//func parseItaExecutionReportBDI(url string) (testSuite string,  environment string,exeHost string, screenDef []string ) {
//	doc, err := xmlquery.LoadURL(url)
//	if err != nil {
//		fmt.Println("Error loading the xml report")
//		return
//	}
//
//	generalNode := xmlquery.Find(doc, "//General")
//	exeHost = generalNode[0].SelectAttr("host")
//
//	listDetailsNode := xmlquery.Find(doc, "//Details")
//	listNodeArgsNode := xmlquery.Find(doc, "//NodeArgs")
//
//
//	for _, node := range listDetailsNode {
//		if strings.Contains(node.InnerText(), "appConfig") {
//			fmt.Println(node.InnerText())
//			break
//		}
//	}
//	for i, node := range listDetailsNode {
//		if strings.Contains(node.InnerText(), "App ID") {
//			var screens []string
//			screens[0] = node.InnerText()
//			screenDef = screens
//		}
//		if listNodeArgsNode[i].SelectAttr("tag") == "Failed" {
//			break
//		}
//	}
//	return testSuite, environment, exeHost, screenDef
//}
//
//func getUrlAndDescription(node *html.Node) (urlList []string, description []string) {
//	descriptionNode := htmlquery.Find(node, `//tr[@class="testRow statusFAILED"]/td[@class="dataCol textCol dataCol2"]`)
//	for _, node := range descriptionNode {
//		description = append(description, renderNode(node))
//	}
//
//	urlListNode := htmlquery.Find(node, `//tr[@class="testRow statusFAILED"]/td[@class="dataCol linkCol dataCol7"]/a`)
//	for _ , node = range urlListNode {
//		for _, attr := range node.Attr {
//			if attr.Key == "href" {
//				urlList = append(urlList, attr.Val)
//			}
//		}
//	}
//	return urlList, description
//}
