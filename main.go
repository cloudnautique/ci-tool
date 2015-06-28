package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ci-tool"
	app.Usage = "Print Test Failures"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "with-skipped",
			Usage: "print skipped tests",
		},
		cli.BoolFlag{
			Name:   "escaped-newline",
			Usage:  "escapes newlines",
			EnvVar: "ESCAPED_NEWLINE",
		},
	}

	app.Action = func(c *cli.Context) {
		report, err := NewTestSuiteReport(c.Args()[0], &TestSuiteReportFormat{
			WithSkipped:    c.Bool("with-skipped"),
			NewLineEscaped: c.Bool("escaped-newline"),
		})
		if err != nil {
			log.Fatalf("Could not parse, Exiting")
		}
		report.Print()
	}

	app.Run(os.Args)
}
