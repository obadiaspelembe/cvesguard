package utils

import (
	"fmt"
)

type VulnerabilityCount struct {
	Low      int `json:"low"`
	Medium   int `json:"medium"`
	High     int `json:"high"`
	Critical int `json:"critical"`
}

func ExecPolicy(policyPath string, reportPath string) bool {
	policy := readPolicyFile(policyPath)
	report := getReport(reportPath)

	if *policy.Spec.Config.Vulnerability.Critical > report.Critical {
		fmt.Println("Critical unexpected ")
		return false
	}

	if *policy.Spec.Config.Vulnerability.High > report.High {
		fmt.Println("High unexpected ")
		return false
	}

	if *policy.Spec.Config.Vulnerability.Medium > report.Medium {
		fmt.Println("Medium unexpected ")
		return false
	}

	if *policy.Spec.Config.Vulnerability.Low > report.Low {
		fmt.Println("Low unexpected ")
		return false
	}

	return true
}

func getReport(path string) VulnerabilityCount {

	report := readReportFile(path)

	var vCount VulnerabilityCount

	for _, vRuns := range *report.Runs {
		for _, vCounts := range *vRuns.Results {
			vCountBuffer := generateVulnerabilityCount(*vCounts.Message.Text)

			vCount.Critical = vCount.Critical + vCountBuffer.Critical
			vCount.High = vCount.High + vCountBuffer.High
			vCount.Medium = vCount.Medium + vCountBuffer.Medium
			vCount.Low = vCount.Low + vCountBuffer.Low
		}
	}

	return vCount
}
