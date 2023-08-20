package utils

func ExecPolicy(policyPath string, reportPath string) bool {
	policy := readPolicyFile(policyPath)
	report := getReport(reportPath)

	skipCheck := checkPackagePolicy(report.Packages, policy)

	if *policy.Spec.Config.Vulnerability.Critical <= report.Critical {
		logVulnerability("Critical", report.Critical, *policy.Spec.Config.Vulnerability.Critical)
		skipCheck = true
	}

	if *policy.Spec.Config.Vulnerability.High <= report.High {
		logVulnerability("High", report.High, *policy.Spec.Config.Vulnerability.High)
		skipCheck = true
	}

	if *policy.Spec.Config.Vulnerability.Medium <= report.Medium {
		logVulnerability("Medium", report.Medium, *policy.Spec.Config.Vulnerability.Medium)
		skipCheck = true
	}

	if *policy.Spec.Config.Vulnerability.Low <= report.Low {
		logVulnerability("Low", report.Low, *policy.Spec.Config.Vulnerability.Low)
		skipCheck = true
	}

	return skipCheck
}

func checkPackagePolicy(cvesPackages []CVESData, policy Policy) bool {

	skipCheck := false

	for _, pPackage := range policy.Spec.Config.Packages {
		for _, cvesPackage := range cvesPackages {
			skip := checkPackageAction(cvesPackage, *pPackage)

			if skip {
				skipCheck = skip
				break
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
