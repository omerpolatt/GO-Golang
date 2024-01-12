package main

import (
	"fmt"
	"os"
	"strconv"
)

func power(taban, üst int) int {

	carpim := 1
	if üst < 0 {
		return 1
	}

	for i := 1; i <= üst; i++ {

		carpim = carpim * taban

	}
	return carpim
}

func main() {

	args := os.Args

	bir, err1 := strconv.Atoi(args[1])
	if err1 != nil {
		fmt.Println("hata", err1)
	}
	iki, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		fmt.Println("hata", err1)
	}

	fmt.Println(power(bir, iki))
}
