package sanitize

import "strings"

// LogValue sanitizes untrusted user input for using it in logs.
func LogValue(userInput string) string {
	cleanInput := strings.ReplaceAll(userInput, "\n", "")
	cleanInput = strings.ReplaceAll(cleanInput, "\r", "")
	return cleanInput
}
