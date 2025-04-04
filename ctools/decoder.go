package ctools

import (
	"errors"
	"strconv"
	"strings"
)

// Handle decoding submitted content/text includes square brackets and may include ASCII art
func DecodeInput(input string) (string, error) {
	var artstr strings.Builder // concatenation strings inside brackets
	var output strings.Builder // whole strings output
	var isInside bool
	var decoded string
	var err error

	if !BalancedBrackets(input) {
		return "", errors.New("square brackets are unbalanced") // Validate bracket balance
	}

	for _, char := range input {
		if char == '[' { // search the first bracket
			isInside = true
		}
		if isInside {
			artstr.WriteRune(char) // adding characters in to artstr
			if char == ']' {       // searching for closing parenthesis
				isInside = false
				decoded, err = ToArt(artstr.String()) // Decode the content within brackets
				if err != nil {
					return "", err
				}
				output.WriteString(decoded)
				artstr.Reset()
			}
		} else {
			output.WriteRune(char) // Write characters outside of brackets directly to output
		}
	}
	return output.String(), nil
}

// decodes a string inside square brackets
func ToArt(input string) (string, error) {
	split := strings.SplitN(input, " ", 2)
	if len(split) != 2 {
		return "", errors.New("arguments are not separated by a space")
	}

	number := split[0]
	str := split[1]

	// removing parentheses from the beginning and end of a string to get the string inside them
	number = strings.TrimPrefix(number, "[")
	str = strings.TrimSuffix(str, "]")

	if str == "" {
		return "", errors.New("there is no second argument")
	}
	// checking if a string contains square brackets
	if strings.ContainsAny(str, "[]") {
		return "", errors.New("square brackets are not printable")
	}

	num, err := strconv.Atoi(number)
	if err != nil {
		return "", errors.New("the first argument is not a number")
	}
	return strings.Repeat(str, num), nil
}

// BalancedBrackets checks if the string has balanced square brackets
func BalancedBrackets(input string) bool {
	var count int
	for _, char := range input {
		switch char {
		case '[':
			count++
		case ']':
			count--
		}
		if count < 0 {
			return false // Early exit if closing bracket comes before a matching opening bracket
		}
	}
	return count == 0
}
