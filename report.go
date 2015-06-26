package main

import (
	"fmt"
)

type TestSuiteReport struct {
	Skipped []JUnitTestCase
	Failed  []JUnitTestCase
}

func GenTestSuiteReport(testcases []JUnitTestCase) TestSuiteReport {
	var failures []JUnitTestCase
	var skipped []JUnitTestCase

	for _, testCase := range testcases {
		if testCase.SkipMessage != nil {
			skipped = append(skipped, testCase)
		}
		if testCase.Failure != nil {
			failures = append(failures, testCase)
		}
	}

	report := TestSuiteReport{
		Failed:  failures,
		Skipped: skipped,
	}

	return report
}

func (report TestSuiteReport) Print() {
	for _, failure := range report.Failed {
		fmt.Printf("Failure: %v\n", failure.Name)
	}
}
