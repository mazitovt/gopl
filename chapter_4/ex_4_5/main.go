package main

import "fmt"

func EliminateDuplicates(input []string) []string {
	for i := 0; i < len(input)-1; {
		if input[i] == input[i+1] {
			input = append(input[:i+1], input[i+2:]...)
		} else {
			i++
		}
	}

	return input
}

func main() {
	ar := []string{"1", "2", "2", "3", "4", "4", "4", "5"}
	fmt.Printf("%q", EliminateDuplicates(ar))
}
