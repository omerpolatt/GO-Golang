package main

import (
	"fmt"
	"os"
)

func BasicAtoi2(s string) int {
	i := 0
	for _, c := range s {
		if !(c >= '0' && c <= '9') { // rakam değilse 0 döndür mantığı burdaki
			return 0 // fonksiyon sonlanıyor herhangi başka bir tipde değer görünce
		} else {
			i = i*10 + int(c-'0') // basamak olayı  ve ascii den  kurtarma kısmı
			// 12345 de 123 aldı 4 ü laması için 1240 oluyor üzerine 4 ekliyor
		}

	}
	return i
}

func main() {
	args := os.Args

	if len(args) == 2 {
		result := BasicAtoi2(args[1])
		fmt.Println(result)

	} else {
		fmt.Println("Hatali argüman girişi ")
	}

}
