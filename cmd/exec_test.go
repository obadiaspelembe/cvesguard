package cmd

import (
	"obadiaspelembe/cves-guard/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	result := utils.ExecPolicy("policy.example.yaml", "cves-report.example.json")
	assert.Equal(t, true, result, "Should return true")
}
