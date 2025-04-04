package ctools

import (
	"errors"
	"fmt"
	"strings"
)

// Encodes given ASCII art to compressed string format
func EncodeFromArt(input string) (string, error) {
	var result strings.Builder

	count := 1  //if 1 char repeats
	count1 := 1 // if 2 chars repeat

	if len(input) == 0 {
		return "", errors.New("input is empty")
	}

	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && runes[i] == runes[i+1] { // Checks for duplicate characters
			count++
		} else {
			if count > 1 {
				result.WriteString(fmt.Sprintf("[%d %c]", count, runes[i])) // Writing duplicate characters to the output
				count = 1
			} else {
				if i+3 < len(runes) && runes[i] == runes[i+2] && runes[i+1] == runes[i+3] { // Checks for duplicate pairs of characters
					count1++
					i++
				} else {
					if count1 > 1 { //
						result.WriteString(fmt.Sprintf("[%d %s]", count1, string(runes[i:i+2])))
						count1 = 1
						i++
					} else { // if no duplicates, write the character to the output stream as is
						result.WriteRune(runes[i])
					}
				}
			}
		}
	}
	return result.String(), nil
}
