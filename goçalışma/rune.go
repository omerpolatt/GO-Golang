package main

import (
	"fmt"
	"os"
	"strconv"
)

func first_rune(str string) rune {
	dizi := []rune(str)
	if len(dizi) > 0 {
		for range dizi {
			return dizi[0]

		}

	}

	return 0
}

func n_rune(str string, n int) rune {

	dizi := []rune(str)
	uzunluk := len(str)
	for i := 0; i <= uzunluk-1; i++ {
		if i == n {
			return dizi[i]
		}
	}
	return 0
}

func last_rune(str string) rune {

	dizi := []rune(str)

	uzunluk := len(dizi)

	if uzunluk > 0 {

		return dizi[uzunluk-1]

	}

	return 0

}

func main() {

	args := os.Args[1:]

	if len(args) == 1 {

		fmt.Printf("%c\n", last_rune(args[0]))
		fmt.Printf("%c\n", first_rune(args[0]))

	}

	if len(args) == 2 {

		cevir, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("Hata", err)
		}
		if len(args) == 2 {

			fmt.Printf("%c", n_rune(args[0], cevir))

		}

	}

}
