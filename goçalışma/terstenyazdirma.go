package main

import (
	"fmt"
	"os"
)

func terstenyazdir(s string) {
	e_sayi := len(s)

	yazi := make([]rune, e_sayi)
	for i := e_sayi - 1; i >= 0; i-- {
		yazi = append(yazi, rune(s[i]))
	}

	yazi = yazi[e_sayi:]

	fmt.Println(string(yazi))
}

func main() {
	args := os.Args

	if len(args) == 2 {

		terstenyazdir(args[1])
	}

}
