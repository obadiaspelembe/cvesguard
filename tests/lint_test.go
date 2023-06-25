package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"obadiaspelembe/cves-guard/utils"
)

func TestLintValidate(t *testing.T) {
	result := utils.Validate("policy.example.yaml", "cves-report.example.json")
	assert.Equal(t, false, result, "Should return false")
}