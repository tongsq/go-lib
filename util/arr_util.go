package util

func InArrayStr(needle string, stack []string) bool {
	for _, val := range stack {
		if val == needle {
			return true
		}
	}
	return false
}

func InArrayInt(needle int, stack []int) bool {
	for _, val := range stack {
		if val == needle {
			return true
		}
	}
	return false
}

func InArrayInt32(needle int32, stack []int32) bool {
	for _, val := range stack {
		if val == needle {
			return true
		}
	}
	return false
}

func InArrayInt64(needle int64, stack []int64) bool {
	for _, val := range stack {
		if val == needle {
			return true
		}
	}
	return false
}
