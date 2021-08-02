package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq(file string) map[string]int {
	words := make(map[string]int)

	if f, err := os.OpenFile(file, os.O_RDONLY, 0444); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words[scanner.Text()]++
		}
	}

	return words
}

func main() {
	for k, v := range wordfreq("test.txt") {
		fmt.Printf("%10s: %d\n", k, v)
	}
}
