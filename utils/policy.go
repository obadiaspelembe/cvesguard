package utils

import "fmt"

type VulnerabilityCount struct {
	Low      int `json:"low"`
	Medium   int `json:"medium"`
	High     int `json:"high"`
	Critical int `json:"critical"`
}

func ApplyPolicy(policyPath string, reportPath string) {
	policy := readPolicyFile(policyPath)
	report := getReport(reportPath)

	if policy.Spec.Config.Low > report.Low {
		fmt.Println("Low unexpected ")
	}
	if policy.Spec.Config.Medium > report.Medium {
		fmt.Println("Medium unexpected ")
	}
	if policy.Spec.Config.High > report.High {
		fmt.Println("High unexpected ")
	}
	if policy.Spec.Config.Critical > report.Critical {
		fmt.Println("Critical unexpected ")
	}
}

func getReport(path string) VulnerabilityCount {

	report := readReportFile(path)

	var vCount VulnerabilityCount

	for _, vRuns := range report.Runs {
		for _, vCounts := range vRuns.Tool.Driver.Rules {
			if vCounts.Properties.Tags[0] == "LOW" || vCounts.Properties.Tags[0] == "UNSPECIFIED" {
				vCount.Low += 1
			}
			if vCounts.Properties.Tags[0] == "MEDIUM" {
				vCount.Medium += 1
			}
			if vCounts.Properties.Tags[0] == "HIGH" {
				vCount.High += 1
			}
			if vCounts.Properties.Tags[0] == "CRITICAL" {
				vCount.Critical += 1
			}
		}
	}

	return vCount
}
