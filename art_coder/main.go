package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	encode := flag.Bool("e", false, "Encode the ASCII art")
	multiline := flag.Bool("m", false, "Use for multi-line input or txt file")
	flag.Parse()

	if len(flag.Args()) == 0 {
		PrintUsage()
		return
	}

	args := flag.Args()
	input := strings.Join(args, " ")

	var content string
	var err error

	if *multiline {
		content, err = ReadFile(input)
		if err != nil {
			log.Fatalf("Failed to read input file: %v", err)
			return
		}
	} else if *encode {
		content, err = ReadFile(input)
		if err != nil {
			content = input // if file is unavialable for reading, input supposed to be strings
		}
	} else {
		content = input
	}

	if err := ValidateInput(*encode, content); err != nil {
		HandleError(err.Error())
		return
	}

	if *encode {
		result := encodeFromArt(content)
		fmt.Println("Encoded content:")
		fmt.Println(result)
	} else {
		result := decodeInput(content)
		fmt.Println("Decoded content:")
		fmt.Println(result)
		//} else {
		//fmt.Println(encodeFromArt(content))
		//PrintUsage()
	}
}
