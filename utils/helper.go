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

func generateVulnerabilityCount(text string) VulnerabilityCount {

	textList := strings.Split(text, "\n")

	var vCount VulnerabilityCount

	for _, textItem := range textList {

		bufferItem := strings.Split(textItem, ":")

		if len(bufferItem) > 1 {

			bufferItemStrIndex := strings.ReplaceAll(bufferItem[0], " ", "")
			bufferItemStrValue := strings.ReplaceAll(bufferItem[1], " ", "")

			if bufferItemStrIndex == VULNERABILITY_SEVERITY {

				if bufferItemStrValue == SERVERITY_LOW || bufferItemStrValue == SERVERITY_UNSPECIFIED {
					vCount.Low += 1
				}
				if bufferItemStrValue == SERVERITY_MEDIUM {
					vCount.Medium += 1
				}
				if bufferItemStrValue == SERVERITY_HIGH {
					vCount.High += 1
				}
				if bufferItemStrValue == SERVERITY_CRITICAL {
					vCount.Critical += 1
				}
			}
		}
	}

	return vCount
}
