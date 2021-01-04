package util

func MapCopy(m map[string]interface{}) map[string]interface{} {
	n := make(map[string]interface{})
	for k, v := range m {
		n[k] = v
	}
	return n
}
