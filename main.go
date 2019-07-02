package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"os"
	"strings"
)

// Output is the output format for the parsed data
type Output struct {
	Status   bool     `json:"status"`
	Checks   int      `json:"checks"`
	Failed   []Result `json:"failed"`
	Warnings []string `json:"warnings"`
}

// Result holds the information about an individual check
type Result struct {
	Title  string `json:"title"`
	Rule   string `json:"rule"`
	Result string `json:"result"`
}

const version = "0.1.0"

var (
	flagVersion = flag.Bool("version", false, "Display version information")
)

func init() {
	flag.Parse()
}

func main() {
	if *flagVersion {
		fmt.Println("oscap-json", version)
		return
	}

	o := Parse(os.Stdin)

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(o); err != nil {
		fmt.Println(err)
	}
}

// Parse reads from in and converts it into an output object
func Parse(in io.Reader) Output {
	scanner := bufio.NewScanner(bufio.NewReader(in))

	processedOutput := Output{
		Status: true,
		Failed: []Result{},
	}

	res := Result{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "WARNING") {
			processedOutput.Warnings = append(processedOutput.Warnings, line)
			continue
		}

		if strings.HasPrefix(line, "Title") {
			res.Title = strings.TrimSpace(strings.TrimPrefix(line, "Title"))
			processedOutput.Checks = processedOutput.Checks + 1
		}

		if strings.HasPrefix(line, "Rule") {
			res.Rule = strings.TrimSpace(strings.TrimPrefix(line, "Rule"))
		}

		if strings.HasPrefix(line, "Result") {
			res.Result = strings.TrimSpace(strings.TrimPrefix(line, "Result"))

			if isFailed(res.Result) {
				processedOutput.Status = false
				processedOutput.Failed = append(processedOutput.Failed, res)
			}

			res = Result{}
		}
	}

	return processedOutput
}

func isFailed(s string) bool {
	var passResults = []string{
		"pass",
		"skipped",
		"notchecked",
	}

	for _, r := range passResults {
		if s == r {
			return false
		}
	}

	return true
}
