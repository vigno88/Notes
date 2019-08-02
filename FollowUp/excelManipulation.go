package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func parseExcelFollowUp() {
	list := []Report{}

	excelFileName := "followUpV3.xlsx"
	excelFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error occured when opening the excel file")
	}
	sheet := excelFile.Sheets[1]
	for _, row := range sheet.Rows {
		newReport := Report{}
		if row != nil {
			newReport.TestSuite = row.Cells[0].String()
			newReport.TestFailing = row.Cells[1].String()
			newReport.Branch = row.Cells[2].String()
			newReport.Date = row.Cells[3].String()
			newReport.Environment = row.Cells[4].String()
			newReport.Rerun = row.Cells[5].String()
			newReport.SameError = row.Cells[6].String()
			newReport.ReportUrl= row.Cells[7].String()
			newReport.ExeHost = row.Cells[8].String()
			screen := []string{}
			screen = append(screen, row.Cells[9].String())
			newReport.ScreenDef = screen
			newReport.FailSince = row.Cells[10].String()
			newReport.Error = row.Cells[11].String()
			newReport.Details = row.Cells[12].String()
			newReport.PassRelaunch = row.Cells[13].String()
			newReport.ActionTaken = row.Cells[14].String()
			list = append(list, newReport)
		}
	}
	reports = list
}

func addRowExcelFollowUp(report Report) {
	excelFileName := "followUpV3.xlsx"
	excelFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error occured when opening the excel file")
	}
	sheet := excelFile.Sheets[1]
	newRow := sheet.AddRow()

	cell1 := newRow.AddCell()
	cell1.Value = report.TestSuite
	cell2 := newRow.AddCell()
	cell2.Value = report.TestFailing
	cell3 := newRow.AddCell()
	cell3.Value = report.Branch
	cell4 := newRow.AddCell()
	cell4.Value = report.Date
	cell5 := newRow.AddCell()
	cell5.Value = report.Environment
	cell9 := newRow.AddCell()
	cell9.Value = report.Rerun
	cell10 := newRow.AddCell()
	cell10.Value = report.SameError
	cell6 := newRow.AddCell()
	cell6.Value = report.ReportUrl
	cell7 := newRow.AddCell()
	cell7.Value = report.ExeHost
	cell8 := newRow.AddCell()
	cell8.Value = report.ScreenDef[0]
	cell11 := newRow.AddCell()
	cell11.Value = report.FailSince
	cell12 := newRow.AddCell()
	cell12.Value = report.Error
	cell13 := newRow.AddCell()
	cell13.Value = report.Details
	cell14 := newRow.AddCell()
	cell14.Value = report.PassRelaunch
	cell15 := newRow.AddCell()
	cell15.Value = report.ActionTaken

	err = excelFile.Save(excelFileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func queryExcel(contains string, filter string) (report []Report) {
	var query []Report
	for _ , item := range reports {
		switch filter {
		case "Test Suite":
			if strings.Contains(item.TestSuite,contains) {
				query = append(query, item)
			}
		case "Test failling":
			if strings.Contains(item.TestFailing,contains) {
				query = append(query, item)
			}
		case "Branch":
			if strings.Contains(item.Branch,contains) {
				query = append(query, item)
			}
		case "Date":
			if strings.Contains(item.Date,contains) {
				query = append(query, item)
			}
		case "Environment":
			if strings.Contains(item.Environment,contains) {
				query = append(query, item)
			}
		case "Report Url":
			if strings.Contains(item.ReportUrl,contains) {
				query = append(query, item)
			}
		case "Host":
			if strings.Contains(item.ExeHost,contains) {
				query = append(query, item)
			}
		case "ScreenDef":
			if strings.Contains(item.ScreenDef[0],contains) {
				query = append(query, item)
			}
		case "Error":
			if strings.Contains(item.Error,contains) {
				query = append(query, item)
			}
		case "Details":
			if strings.Contains(item.Details,contains) {
				query = append(query, item)
			}
		}
	}
	return query
}


