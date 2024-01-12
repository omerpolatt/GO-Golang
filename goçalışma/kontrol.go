package main

import (
	"fmt"
	"os"
)

func IsUpper(s string) bool {

	harfler := []rune(s)
	boyut := len(harfler)

	for i := 0; i < boyut; i++ {

		if harfler[i] < 'A' || harfler[i] > 'Z' {
			return false
		}

	}
	return true
}

func main() {

	args := os.Args[1:]

	fmt.Println(IsUpper(args[0]))
}
