package main

import (
	"fmt"
	"os"
)

func alfacount(str string) int {

	sayac := 0
	dizi := []rune(str)
	uzunluk := len(dizi)

	for i := 0; i <= uzunluk-1; i++ {

		if ('a' <= dizi[i] && dizi[i] <= 'z') || ('A' <= dizi[i] && dizi[i] <= 'Z') {
			sayac++
		}
	}
	return sayac

}

func main() {

	args := os.Args[1:]

	fmt.Println(alfacount(args[0]))
}
