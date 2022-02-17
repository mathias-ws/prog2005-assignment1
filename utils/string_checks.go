package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// CheckIfStringIsNotEmpty checks if a string contains only letters.
func CheckIfStringIsNotEmpty(str string) bool {
	//Removes whitespaces
	strToCheck := strings.ReplaceAll(str, " ", "")

	if strToCheck != "" {
		fmt.Println(strToCheck)
		fmt.Println(str)
		return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(strToCheck)
	}
	return false
}
