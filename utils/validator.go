package utils

import (
	"fmt"
	"io/ioutil"

	"reflect"

	"gopkg.in/yaml.v3"
)

const POLICY_VERSION = "PolicyVersion"
const POLICY = "Policy"
const NAME = "Name"

type PolicyItem struct {
	Type   string `json:"type" binding:"required"`
	Target string `json:"target" binding:"required"`
}

type Policy struct {
	PolicyVersion string       `json:"policyVersion"`
	Name          string       `json:"name"`
	Policy        []PolicyItem `json:"policy"`
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
		PolicyVersion: config.PolicyVersion,
		Name:          config.Name,
		Policy:        config.Policy,
	}

}

func Validate(policyFile string) Policy {
	policyJson := readPolicyFile(policyFile)

	// Check if the policyVersion is Zero
	if reflect.ValueOf(policyJson).FieldByName(POLICY).IsZero() {
		fmt.Printf("%s field is invalid \n", POLICY)
	}

	if reflect.ValueOf(policyJson).FieldByName(POLICY_VERSION).IsZero() {
		fmt.Printf("%s field is invalid \n", POLICY_VERSION)
		fmt.Println(policyJson.PolicyVersion)
	}

	if reflect.ValueOf(policyJson).FieldByName(NAME).IsZero() {
		fmt.Printf("%s field is invalid \n", NAME)
	}

	return policyJson
}
