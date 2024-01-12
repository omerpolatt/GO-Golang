package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Printstr(s string) {

	for _, i := range s {
		z01.PrintRune(i)

	}
	z01.PrintRune('\n')
}

func main() {

	argumanlar := os.Args

	if len(argumanlar) == 2 {
		Printstr(argumanlar[1])
	} else {
		fmt.Println("geçerli argüman sayisinda değer gir ")
	}
}
