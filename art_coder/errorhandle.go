package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ValidateInput verifies input for errors
func ValidateInput(isEncoding bool, inputText string) error {
	if strings.TrimSpace(inputText) == "" {
		return fmt.Errorf("the input string is empty")
	}
	if isEncoding {
		return nil
	}
	if !balancedBrackets(inputText) {
		return fmt.Errorf("square brackets marks are unbalanced")
	}

	// Checking decode format
	if !isValidDecodeFormat(inputText) {
		return fmt.Errorf("the first argument is not a number")
	}

	return nil
}

// balancedBrackets checks brackets balancing
func balancedBrackets(input string) bool {
	var count int
	for _, char := range input {
		switch char {
		case '[':
			count++
		case ']':
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// isValidDecodeFormat verifies correct input format
func isValidDecodeFormat(input string) bool {
	var inBracket bool
	var numberFound bool

	for _, char := range input {
		switch char {
		case '[':
			if inBracket {
				return false // Nested brackets are not allowed
			}
			inBracket = true
			numberFound = false
		case ']':
			if !inBracket {
				return false // Closing bracket without opening bracket
			}
			inBracket = false
		case ' ':
			if inBracket && !numberFound {
				return false // Space found inside brackets before number
			}
		default:
			if inBracket && !numberFound {
				if _, err := strconv.Atoi(string(char)); err != nil {
					return false // The first argument is not a number
				}
				numberFound = true
			}
		}
	}
	return !inBracket
}

// HandleError checks for errors and prints an error message if one occurs
func HandleError(message string) {
	fmt.Println("Error:")
	fmt.Println(message)
	os.Exit(1) // exit from program if error occures
}

func PrintUsage() {
	fmt.Println("Usage if strings: \ngo run .'<decode-content>'.\ngo run . -e '<encode-content>'.")
	fmt.Println("\nUsage if .txt file: \ngo run . [-m|-e] <strings-or-art.txt>")
	os.Exit(0)
}
