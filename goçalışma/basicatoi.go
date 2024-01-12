package main

import (
	"fmt"
	"os"
)

func BasicAtoi(s string) int {
	sayi := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
			sayi = sayi*10 + int(i-'0')
		}
	}
	return sayi
}

// return kullandıysan  fonk da yazdırmaak için main içinde yazdırma fonksiyonu kullan

func main() {
	args := os.Args

	if len(args) == 2 {
		result := BasicAtoi(args[1])
		fmt.Println(result)
	} else {
		fmt.Println("Hatali argüman girişi")
	}
}
