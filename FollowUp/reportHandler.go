package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type testId struct {
	Id int `json:"id"`
}

type Report struct {
	TestSuite   string   `json:"testSuite"`
	TestFailing string   `json:"testFailing"`
	Branch      string   `json:"branch"`
	Date        string   `json:"date"`
	Environment string   `json:"environment"`
	ReportUrl   string   `json:"reportUrl"`
	ExeHost     string   `json:"exeHost"`
	ScreenDef   []string `json:"screenDef"`

	Rerun        string `json:"rerun"`
	SameError    string `json:"sameError"`
	FailSince    string `json:"failSince"`
	Error        string `json:"error"`
	Details      string `json:"details"`
	PassRelaunch string `json:"passRelaunch"`
	ActionTaken  string `json:"actionTaken"`
}

type Url struct {
	UrlOfReport string `json:"urlOfReport"`
}

var reports []Report
var newReports []Report
var parsed bool
var listFailedUrl []Url

// Post - /reports
func createReportsHandler(w http.ResponseWriter, r *http.Request) {
	//parse the url form of the UI
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := r.Form.Get("url")
	url = strings.Trim(url, " ")
	url += "ItaFlowReport.html"
	urlList = parseItaFlowReport(url)
	parseItaExecutionReport()
	newReports = fillReports()
	fmt.Println(newReports)
	err = json.NewEncoder(w).Encode(newReports)
	if err != nil {
		fmt.Println(err)
	}
	parseExcelFollowUp()
	parsed = true
}

// Get - /report
func getReportsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Request")
	fmt.Println(newReports)
	err := json.NewEncoder(w).Encode(newReports)
	if err != nil {
		fmt.Println(err)
	}
}

func saveReport(w http.ResponseWriter, r *http.Request) {

	fmt.Println("save Rport")
	var reportGet Report
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reportGet)
	if err != nil {
		fmt.Println(err)
	}

	indexGet := findIndex(reportGet)
	report := newReports[indexGet]
	if reportGet.Rerun == "true" {
		report.Rerun = "YES"
	} else {
		report.Rerun = "NO"
	}

	if reportGet.SameError == "true" {
		report.SameError = "YES"
	} else {
		report.SameError = "NO"
	}

	if reportGet.PassRelaunch == "true" {
		report.PassRelaunch = "YES"
	} else {
		report.PassRelaunch = "NO"
	}

	report.ScreenDef[0] = reportGet.ScreenDef[0]
	report.FailSince = reportGet.FailSince
	report.Error = reportGet.Error
	report.Details = reportGet.Details
	report.ActionTaken = reportGet.ActionTaken


	reports = append(reports,report)
	index := findIndex(report)
	if len(newReports) == 1 {
		newReports = nil
	} else {
		newReports = append(newReports[:index], newReports[index+1:]...)
	}
	fmt.Println(report)
	addRowExcelFollowUp(report)
}

func findIndex(report Report) int {
	var index int
	for i ,element := range newReports {
		if element.TestFailing == report.TestFailing {
			index = i
		}
	}
	return index
}

func queryReport(w http.ResponseWriter, r *http.Request) {
	if !parsed {
		parseExcelFollowUp()
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	query := queryExcel(r.Form["contains"][0], r.Form["filter"][0])
	fmt.Println(query)
	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete report")

	var  report Report
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&report)
	if err != nil {
		fmt.Println(err)
	}
	index := findIndex(report)
	if len(newReports) == 1 {
		newReports = nil
	} else {
		newReports = append(newReports[:index], newReports[index+1:]...)
	}

	fmt.Println(newReports)

}