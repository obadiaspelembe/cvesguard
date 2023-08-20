package utils

import (
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkReturnBool(e error) bool {
	if e != nil {
		return false
	}
	return true
}

func generateCVESData(text string) CVESData {

	textList := strings.Split(text, "\n")

	var cvesData CVESData

	for _, textItem := range textList {

		bufferItem := strings.Split(textItem, ":")

		if len(bufferItem) > 1 {

			bufferItemStrIndex := strings.ReplaceAll(bufferItem[0], " ", "")
			bufferItemStrValue := strings.ReplaceAll(bufferItem[1], " ", "")

			if bufferItemStrIndex == VULNERABILITY_SEVERITY {
				cvesData.Severity = bufferItemStrValue
			}

			if bufferItemStrIndex == VULNERABILITY_PACKAGE {
				cvesData.Package = bufferItemStrValue
			}
		}
	}
 
	return cvesData
}

func generateVulnerabilityCount(cvesData CVESData) VulnerabilityCount {

	var vCount VulnerabilityCount

	if cvesData.Severity == SERVERITY_LOW || cvesData.Severity == SERVERITY_UNSPECIFIED {
		vCount.Low += 1
	}
	if cvesData.Severity == SERVERITY_MEDIUM {
		vCount.Medium += 1
	}
	if cvesData.Severity == SERVERITY_HIGH {
		vCount.High += 1
	}
	if cvesData.Severity == SERVERITY_CRITICAL {
		vCount.Critical += 1
	}

	return vCount
}
