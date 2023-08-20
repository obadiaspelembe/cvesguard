package utils

import (
	"fmt"
	"strings"
)

func ExecPolicy(policyPath string, reportPath string) bool {
	policy := readPolicyFile(policyPath)
	report := getReport(reportPath)

	checkPackagePolicy(report.Packages, policy)

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

func checkPackagePolicy(cvesPackages []CVESData, policy Policy) bool {

	skipCheck := true

	for _, pPackage := range policy.Spec.Config.Packages {
		for _, cvesPackage := range cvesPackages {
			 if strings.Contains(cvesPackage.Package, *pPackage.Name) && containerStringInList(*pPackage.Severity, cvesPackage.Severity) {
				fmt.Println(*pPackage.Action)
			}
		}
	}

	return skipCheck
}

func getReport(path string) VulnerabilityCount {

	report := readReportFile(path)

	var vCount VulnerabilityCount

	for _, vRuns := range *report.Runs {
		for _, vCounts := range *vRuns.Results {
			cvesData := generateCVESData(*vCounts.Message.Text)
			vCountBuffer := generateVulnerabilityCount(cvesData)

			vCount.Critical = vCount.Critical + vCountBuffer.Critical
			vCount.High = vCount.High + vCountBuffer.High
			vCount.Medium = vCount.Medium + vCountBuffer.Medium
			vCount.Low = vCount.Low + vCountBuffer.Low
			vCount.Packages = append(vCount.Packages, cvesData)
		}
	}

	return vCount
}
