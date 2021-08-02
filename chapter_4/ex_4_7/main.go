package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(input []byte) {
	swap := func(i, j, l int) {
		for beg, end := i, j; beg < l; beg, end = beg+1, end-1 {
			input[beg], input[end] = input[end], input[beg]
		}
	}
	l := len(input)
	for i := 0; i < l; {
		_, s := utf8.DecodeRune(input[i:])
		swap(i, i+s-1, i+len(input[i:i+s])/2)
		i += s
	}
	swap(0, l-1, l/2)
}

func main() {
	in := []byte("Hello 世界")
	out := []byte("界世 olleH")
	fmt.Println(in)
	reverse(in)
	fmt.Println(in)
	fmt.Println(out)

	in = []byte("ABCD  世界 界界世世 A")
	out = []byte("A 世世界界 界世  DCBA")

	fmt.Println(in)
	reverse(in)
	fmt.Println(in)
	fmt.Println(out)
}
