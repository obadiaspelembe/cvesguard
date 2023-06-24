package utils

import (
	"fmt"

	"reflect"
)

func checkPolicy(policyFile string) bool {

	policyJson := readPolicyFile(policyFile)

	if reflect.ValueOf(policyJson).FieldByName(SPEC).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined.\n", SPEC)
		return false
	}

	if reflect.ValueOf(policyJson).FieldByName(VERSION).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", VERSION)
		return false
	}

	if reflect.ValueOf(policyJson).FieldByName(NAME).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", NAME)
		return false
	}

	return true
}

func checkReport(policyFile string) bool {

	policyJson := readPolicyFile(policyFile)

	if reflect.ValueOf(policyJson).FieldByName(SPEC).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined.\n", SPEC)
		return false
	}

	if reflect.ValueOf(policyJson).FieldByName(VERSION).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", VERSION)
		return false
	}

	if reflect.ValueOf(policyJson).FieldByName(NAME).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", NAME)
		return false
	}

	return true
}

func Validate(policyFile string) {

}
