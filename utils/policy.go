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

func ApplyPolicy(policyPath string, reportPath string) bool {
	policy := readPolicyFile(policyPath)
	report := getReport(reportPath)

	if *policy.Spec.Config.Critical > report.Critical {
		fmt.Println("Critical unexpected ")
		return false
	}

	if *policy.Spec.Config.High > report.High {
		fmt.Println("High unexpected ")
		return false
	}

	if *policy.Spec.Config.Medium > report.Medium {
		fmt.Println("Medium unexpected ")
		return false
	}

	if *policy.Spec.Config.Low > report.Low {
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
			vCountBuffer := getResultProps(*vCounts.Message.Text)

			vCount.Critical = vCount.Critical + vCountBuffer.Critical
			vCount.High = vCount.High + vCountBuffer.High
			vCount.Medium = vCount.Medium + vCountBuffer.Medium
			vCount.Low = vCount.Low + vCountBuffer.Low
		}
	}

	return vCount
}
