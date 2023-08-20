package utils

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)


/*
	Read CVES Policy yaml file and parse it into struct
	:path - policy filename
*/
func readPolicyFile(path string) Policy {

	data, err := os.ReadFile(path)

	var config Policy

	err = yaml.Unmarshal([]byte(data), &config)

	check(err)

	return Policy{
		Version: config.Version,
		Spec:    config.Spec,
	}
}

/*
	Read CVES report json file and parse it into struct
	:path - report filename
*/
func readReportFile(path string) Report {

	data, err := os.ReadFile(path)

	var report Report

	err = json.Unmarshal([]byte(data), &report)

	check(err)

	return Report{
		Version: report.Version,
		Runs:    report.Runs,
	}
}
