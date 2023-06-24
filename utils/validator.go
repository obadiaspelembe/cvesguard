package utils

import (
	"fmt"
	"io/ioutil"

	"reflect"

	"gopkg.in/yaml.v3"
)

const VERSION = "Version"
const POLICY = "Policy"
const NAME = "Name"

type PolicyItem struct {
	Type   string `json:"type" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type Policy struct {
	Version string       `json:"version"`
	Name    string       `json:"name"`
	Policy  []PolicyItem `json:"policy"`
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
		Policy:  config.Policy,
	}

}

func Validate(policyFile string) Policy {

	policyJson := readPolicyFile(policyFile)

	// Check if the policyVersion is Zero
	if reflect.ValueOf(policyJson).FieldByName(POLICY).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined.\n", POLICY)
	
	}

	if reflect.ValueOf(policyJson).FieldByName(VERSION).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", VERSION)
		fmt.Println(policyJson.Version)
	}

	if reflect.ValueOf(policyJson).FieldByName(NAME).IsZero() {
		fmt.Printf("%s field is invalid, check if it's defined. \n", NAME)
	}

	return policyJson
}
