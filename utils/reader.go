package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func readPolicyFile(path string) Policy {

	data, err := ioutil.ReadFile(path)

	var config Policy

	err = yaml.Unmarshal([]byte(data), &config)

	check(err)

	return Policy{
		Version: config.Version,
		Name:    config.Name,
		Spec:    config.Spec,
		Kind:    config.Kind,
	}
}

func readReportFile(path string) Report {

	data, err := ioutil.ReadFile(path)

	var report Report

	err = yaml.Unmarshal([]byte(data), &report)

	check(err)

	return Report{
		Version: report.Version,
		Runs:    report.Runs,
	}
}
