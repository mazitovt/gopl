package main

import (
	"fmt"
	"unicode"
)

func eliminateSpaces(input []byte) []byte {
	for i := 0; i < len(input)-1; {
		if unicode.IsSpace(rune(input[i])) && input[i] == input[i+1] {
			input = append(input[:i+1], input[i+2:]...)
		} else {
			i++
		}
	}

	return input
}

func main() {
	input := []byte("Excuse   me,  can you help   .")
	fmt.Printf("%q", eliminateSpaces(input))
}
