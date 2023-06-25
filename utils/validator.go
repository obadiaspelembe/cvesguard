package utils

import (
	"encoding/json"
	"fmt"
	"github.com/grahms/godantic"
)

func checkPolicy(policyFile string) bool {

	readPolicyFile(policyFile)

	return true
}

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

func Validate(policyFile string, reportFile string) bool {

	condition, err := checkReport(policyFile)

	if !condition {
		fmt.Println("Validation result: ", err)
		return false
	}

	fmt.Println("Configurations are valid")

	return true
}
