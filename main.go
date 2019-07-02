package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"os"
	"strings"
)

type output struct {
	Status   bool     `json:"status"`
	Checks   int      `json:"checks"`
	Failed   []result `json:"failed"`
	Warnings []string `json:"warnings"`
}

type result struct {
	Title  string `json:"title"`
	Rule   string `json:"rule"`
	Result string `json:"result"`
}

func main() {
	o := Parse(os.Stdin)

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(o); err != nil {
		fmt.Println(err)
	}
}

func Parse(in io.Reader) output {
	scanner := bufio.NewScanner(bufio.NewReader(in))

	processedOutput := output{
		Status: true,
		Failed: []result{},
	}

	res := result{}

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

			res = result{}
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
