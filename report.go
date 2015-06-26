package main

import (
	"fmt"
)

type TestSuiteReport struct {
	Skipped []JUnitTestCase
	Failed  []JUnitTestCase
	Errored []JUnitTestCase
}

func GenTestSuiteReport(testcases []JUnitTestCase) TestSuiteReport {
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

	report := TestSuiteReport{
		Failed:  failures,
		Skipped: skipped,
		Errored: errored,
	}

	return report
}

func (report TestSuiteReport) Print(withSkippedFlag bool) {
	printJUnitSlice("Failed", report.Failed)
	printJUnitSlice("Errored", report.Errored)

	if withSkippedFlag {
		printJUnitSlice("Skipped", report.Skipped)
	}
}

func printJUnitSlice(messagePfx string, testcases []JUnitTestCase) {
	for _, testcase := range testcases {
		fmt.Printf("%s: %s\n", messagePfx, testcase.Name)
	}
}
