package utils

import (
	"fmt"
	"strings"
)

func containerStringInList(list []string, target string) bool {
	for _, value := range list {
		if strings.ToUpper(value) == target {
			return true
		}
	}
	return false
}

func checkPackageAction(cvesPackage CVESData, pPackage Package) bool {

	packageAction := strings.ToUpper(*pPackage.Action)

	if packageAction == POLICY_PACKAGE_ACTIONS_DENY {
		pkgs := strings.Join(*pPackage.Severity, ", ")
		pkgs = strings.ToUpper(pkgs)
		fmt.Println("[", packageAction, "]:", " Package ", *pPackage.Name, pkgs, "severity")
		return true
	}

	return false
}

func logVulnerability(severity string, reportSeverity int, policySeverity int) {
	fmt.Println("[ DENY ]: ", severity, "unexpected: Found", reportSeverity, "/", policySeverity)
}
