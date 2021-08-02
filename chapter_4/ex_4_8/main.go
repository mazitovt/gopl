package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func CountCharByCategory(input []byte) map[string]int {
	count := map[string]int{"Digit": 0, "Letter": 0, "Punctuation": 0, "WhiteSpace": 0, "Other": 0}
	for i, s := 0, 0; i < len(input); i += s {
		r, s := utf8.DecodeRune(input[i:])
		switch {
		case unicode.IsDigit(r):
			count["Digit"]++
		case unicode.IsLetter(r):
			count["Letter"]++
		case unicode.IsPunct(r):
			count["Punctuation"]++
		case unicode.IsSpace(r):
			count["WhiteSpace"]++
		default:
			count["Other"]++
		}
		i += s
	}
	return count
}

func main() {
	s := []byte("15Aa界.?\\ >")
	//fmt.Println(CountCharByCategory(s))
	for k, v := range CountCharByCategory(s) {
		fmt.Printf("%15s: %d\n", k, v)
	}
}
