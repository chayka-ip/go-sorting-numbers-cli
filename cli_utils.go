package main

import (
	"regexp"
	"strings"
)

func isYesInput(s string) bool {
	s = getFirstCharFromString(strings.ToLower(s))
	return (s == "\r") || (s == "y")
}

func getNumberFromString(s string) (int, error) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(s, "")
	return tryConvertStringToInt(processedString)
}
