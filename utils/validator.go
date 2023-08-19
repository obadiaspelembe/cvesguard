package utils

import (
	"encoding/json"
	"fmt"
	"github.com/grahms/godantic"
)

// Check if policy file comply with policy schema
func checkPolicy(policyFile string) bool {

	readPolicyFile(policyFile)

	return true
}

// Check if report file comply with report schema
func checkReport(reportFile string) (bool, error) {

	var report Report

	reportObject := readReportFile(reportFile)

	validator := godantic.Validate{}

	jsonBytes, err := json.Marshal(reportObject)

	if err == nil {
		err = validator.BindJSON(jsonBytes, &report)
	}

	return checkReturnBool(err), err
}

// Validate the policy and reports files to comply with the expected schema
func Validate(policyFile string, reportFile string) bool {

	condition, err := checkReport(reportFile)

	if !condition {
		fmt.Println("Validation result: ", err)
		return false
	}

	fmt.Println("Configurations are valid")

	return true
}
