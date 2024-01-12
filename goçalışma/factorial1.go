package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(nb int) int {

	carpim := 1

	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {

			carpim = carpim * i

		}
		return carpim
	}

}

func main() {

	/*arg := 0
	fmt.Println(factorial(arg))*/

	args := os.Args

	input, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	sonuc := factorial(input)
	fmt.Println(sonuc)

}
