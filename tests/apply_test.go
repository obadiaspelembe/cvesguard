package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"obadiaspelembe/cves-guard/utils"
)

func TestApply(t *testing.T) {
	result := utils.ApplyPolicy("policy.example.yaml", "cves-report.example.json")
	assert.Equal(t, true, result, "Should return true")
}