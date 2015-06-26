package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ci-tool"
	app.Usage = "Print Test Failures"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) {
		err := parseFile(c.Args()[0])
		if err != nil {
			log.Fatalf("Could not parse, Exiting")
		}
	}

	app.Run(os.Args)
}

func parseFile(fileName string) error {
	log.Infof("Processing file: %s", fileName)
	junitFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error openining file: %s", fileName)
		return err
	}
	defer junitFile.Close()

	XMLdata, _ := ioutil.ReadAll(junitFile)

	var testSuite JUnitTestSuite
	xml.Unmarshal(XMLdata, &testSuite)

	report := GenTestSuiteReport(testSuite.TestCases)
	report.Print()
	return nil
}