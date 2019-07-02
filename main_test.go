package main_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/manywho/oscap-json"
)

func TestParse(t *testing.T) {
	var testInput = `
Title   Test rule 1
Rule    xccdf_org.ssgproject.content_rule_test_rule_1
Result  pass

Title   Test rule 2
Rule    xccdf_org.ssgproject.content_rule_test_rule_2
Result  fail

Title   Test rule 3
Rule    xccdf_org.ssgproject.content_rule_test_rule_3
WARNING: test warning message
Result  notchecked

Title   Test rule 4
Rule    xccdf_org.ssgproject.content_rule_test_rule_4
Result  skipped`

	reader := bytes.NewBufferString(testInput)

	out := Parse(reader)

	assert.Equal(t, 4, out.Checks)

	assert.Len(t, out.Failed, 1)
	assert.Equal(t, "Test rule 2", out.Failed[0].Title)
	assert.Equal(t, "xccdf_org.ssgproject.content_rule_test_rule_2", out.Failed[0].Rule)
	assert.Equal(t, "fail", out.Failed[0].Result)

	assert.Len(t, out.Warnings, 1)
	assert.Equal(t, "WARNING: test warning message", out.Warnings[0])

}
