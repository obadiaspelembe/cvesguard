package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type PolicyItem struct {
	Type   string `json:"type"`
	Target string `json:"target"`
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

func readFile(path string) {

	data, err := ioutil.ReadFile(path)

	var config Policy
	err = yaml.Unmarshal([]byte(data), &config)

	check(err)

}

func Validate(policy string) {
	readFile(policy)
}
