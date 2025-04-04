package main

import (
	"strconv"
	"strings"
)

// decodes input text using square brackets
func decodeInput(input string) string {
	var artstr strings.Builder // concatenation strings inside brackets
	var output strings.Builder // whole strings output
	var isInside bool
	var decoded string

	if !balancedBrackets(input) {
		HandleError("Square brackets are unbalanced.")
		return ""
	}

	for _, char := range input {
		if char == '[' { // search the first bracket
			isInside = true
		}
		if isInside {
			artstr.WriteRune(char) // adding characters in to artstr
			if char == ']' {       // searching for closing parenthesis
				isInside = false
				decoded = toArt(artstr.String()) // writing decoded text to output
				if decoded == "" {
					return ""
				}
				output.WriteString(decoded)
				artstr.Reset()
			}
		} else {
			output.WriteRune(char)
		}
	}
	return output.String()
}

// decodes a string inside square brackets
func toArt(input string) string {
	split := strings.SplitN(input, " ", 2)
	if len(split) != 2 {
		HandleError("Arguments are not separated by a space.")
		return ""
	}

	number := split[0]
	str := split[1]

	// removing parentheses from the beginning and end of a string to get the string inside them
	number = strings.TrimPrefix(number, "[")
	str = strings.TrimSuffix(str, "]")

	if str == "" {
		HandleError("There is no second argument.")
		return ""
	}
	// checking if a string contains square brackets
	if strings.ContainsAny(str, "[]") {
		HandleError("Square brackets are not printable.")
		return ""
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		HandleError("The first argument is not a number.")
		return ""
	}
	return strings.Repeat(str, num)
}
