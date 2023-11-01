package Service

func GetAddressFromDataSource(userID string) string {
	if userID == "1" {
		return "123 Main St, San Francisco, CA 94101"
	} else if userID == "2" {
		return "456 Main St, San Francisco, CA 94101"
	} else {
		return ""
	}
}
