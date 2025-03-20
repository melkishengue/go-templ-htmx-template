package utils

func GetDefaultString(val string, defaultValue string) string {
	if val == "" {
		return defaultValue
	}

	return val
}

func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
