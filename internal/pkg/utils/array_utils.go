package utils

// Contains = returns true if the string is present in the array, false otherwise
func Contains(arr []string, elem string) bool {
	for _, v := range arr {
		if elem == v {
			return true
		}
	}
	return false
}
