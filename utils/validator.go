package utils

import (
	"fmt"
	"io/ioutil"

	"reflect"

	"gopkg.in/yaml.v3"
)

const VERSION = "Version"
const SPEC = "Spec"
const NAME = "Name"

type SpecItem struct {
	Type   string `json:"type" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type Policy struct {
	Version string     `json:"version"`
	Name    string     `json:"name"`
	Spec    []SpecItem `json:"spec"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readPolicyFile(path string) Policy {

	data, err := ioutil.ReadFile(path)

	var config Policy
	err = yaml.Unmarshal([]byte(data), &config)

	check(err)

	return Policy{
		Version: config.Version,
		Name:    config.Name,
		Spec:    config.Spec,
	}

}

func Validate(policyFile string) Policy {

	policyJson := readPolicyFile(policyFile)

	// Check if the policyVersion is Zero
	if reflect.ValueOf(policyJson).FieldByName(SPEC).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined.\n", SPEC)
	}

	if reflect.ValueOf(policyJson).FieldByName(VERSION).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", VERSION) 
	}

	if reflect.ValueOf(policyJson).FieldByName(NAME).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", NAME)
	}

	return policyJson
}
