package utils

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
