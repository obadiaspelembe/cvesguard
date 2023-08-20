package utils

import (
	"fmt"
	"strings"
)

func containerStringInList(list []string, target string) bool {

	fmt.Println(list, target)
	for _, value := range list {
		if strings.ToUpper(value) == target {
			return true
		}
	}
	return false
}
