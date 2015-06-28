package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

type TestSuiteReport struct {
	Skipped []JUnitTestCase
	Failed  []JUnitTestCase
	Errored []JUnitTestCase
	Format  *TestSuiteReportFormat
}

type TestSuiteReportFormat struct {
	WithSkipped    bool
	NewLineEscaped bool
}

func NewTestSuiteReport(file string, format *TestSuiteReportFormat) (TestSuiteReport, error) {
	var report TestSuiteReport
	parsedJunitFile, err := parseJunitFile(file)
	if err != nil {
		log.Fatalf("Could not create report")
		return report, err
	}

	report.setTestSuiteReport(parsedJunitFile.TestCases)
	report.Format = format

	return report, nil
}

func parseJunitFile(filename string) (JUnitTestSuite, error) {
	var testSuite JUnitTestSuite

	junitFile, err := os.Open(filename)
	if err != nil {
		log.Errorf("Error openining file: %s", filename)
		return testSuite, err
	}
	defer junitFile.Close()

	XMLdata, err := ioutil.ReadAll(junitFile)
	if err != nil {
		log.Errorf("Error parsing Junit file: %s", filename)
		return testSuite, err
	}

	xml.Unmarshal(XMLdata, &testSuite)

	return testSuite, nil
}

func (report *TestSuiteReport) setTestSuiteReport(testcases []JUnitTestCase) {
	var failures []JUnitTestCase
	var skipped []JUnitTestCase
	var errored []JUnitTestCase

	for _, testCase := range testcases {
		if testCase.SkipMessage != nil {
			skipped = append(skipped, testCase)
		}
		if testCase.Failure != nil {
			failures = append(failures, testCase)
		}
		if testCase.Error != nil {
			errored = append(errored, testCase)
		}
	}

	report.Failed = failures
	report.Skipped = skipped
	report.Errored = errored

}

func (report TestSuiteReport) Print() {
	newLine := getNewLine(report.Format.NewLineEscaped)
	printJUnitSlice("Failed", newLine, report.Failed)
	printJUnitSlice("Errored", newLine, report.Errored)

	if report.Format.WithSkipped {
		printJUnitSlice("Skipped", newLine, report.Skipped)
	}
}

func getNewLine(flag bool) string {
	newLine := "\n"
	if flag {
		newLine = "\\n"
	}
	return newLine
}

func printJUnitSlice(messagePfx string, newLine string, testcases []JUnitTestCase) {
	for _, testcase := range testcases {
		fmt.Printf("%s: %s%s", messagePfx, newLine, testcase.Name)
	}
}
