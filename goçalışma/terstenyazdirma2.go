package main

import (
	"fmt"
	"os"
)

func terstenyazdir2(s string) {
	e_sayi := len(s)

	//var yazi string oalrak da tanımlarız 
	yazi := ""
	for i := e_sayi - 1; i >= 0; i-- {
		yazi += string(s[i])
	}

	fmt.Println(string(yazi))
}

func main() {
	args := os.Args

	if len(args) == 2 {

		terstenyazdir2(args[1])
	}

}
