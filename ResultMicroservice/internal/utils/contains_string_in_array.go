package helper

func ContainsString(array []string, target string) bool {
	for _, str := range array {
		if str == target {
			return true
		}
	}
	return false
}
