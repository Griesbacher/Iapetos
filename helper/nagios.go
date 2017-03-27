package helper

import "fmt"

func ReturnCodeToString(returnCode int) string {
	switch returnCode {
	case 0:
		return "OK"
	case 1:
		return "Warning"
	case 2:
		return "Critical"
	case 3:
		return "Unknown"
	default:
		return fmt.Sprintf("Could not convert returncode %d", returnCode)
	}
}
