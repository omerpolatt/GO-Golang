package main

import (
	"fmt"
	"os"
	"strconv"
)

func asal(nb int) bool {
	sayac := 0

	if nb < 2 {
		return false
	} else {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				sayac++
			}
		}

		if sayac == 0 {
			return true
		} else {
			return false
		}
	}
}
func main() {
	args := os.Args

	cevir, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("hATA", err)
	}
	sonuc := asal(cevir)
	fmt.Println(sonuc)
}
