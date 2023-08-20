package cmd

import (
	"obadiaspelembe/cves-guard/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLintValidate(t *testing.T) {
	result := utils.Validate("policy.example.yaml", "cves-report.example.json")
	assert.Equal(t, true, result, "Should return true")
}
