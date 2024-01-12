package main

import (
	"fmt"
	"os"
)

func uppercase(str string) string {

	dizi := []rune(str)
	boyut := len(dizi)
	cikti := ""

	for i := 0; i < boyut; i++ {
		if (dizi[i] >= 'a' && dizi[i] <= 'z') || (dizi[i] >= 'A' && dizi[i] <= 'Z') {
			if dizi[i] >= 'a' && dizi[i] <= 'z' {
				dizi[i] = dizi[i] - 32
			} else {
				dizi[i] = dizi[i]
			}
		}
		cikti = cikti + string(dizi[i])
	}
	return cikti

}

func main() {

	args := os.Args[1:]

	if len(args) == 1 {
		output := uppercase(args[0])
		fmt.Println(output)
	}
}
