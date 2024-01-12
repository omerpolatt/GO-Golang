package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	lastWord := ""

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != ' ' {
			lastWord = string(input[i]) + lastWord
		} else if lastWord != "" {
			break
		}
	}

	if lastWord == "" {
		z01.PrintRune('\n')
		return
	}

	for _, r := range lastWord {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
