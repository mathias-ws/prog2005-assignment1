package utils

import (
	"regexp"
	"strings"
)

// CheckIfStringIsNotEmpty checks if a string contains only letters.
func CheckIfStringIsNotEmpty(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(strToCheck)
	}
	return false
}
