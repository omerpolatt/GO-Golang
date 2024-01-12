package main

import (
	"fmt"
	"os"
)

func ReplaceChars(düzeltme, aranan, degisen string) string {

	kelime := []rune(düzeltme)

	boyut := len(kelime)

	for i := 0; i < boyut; i++ {

		if string(kelime[i]) == string(aranan[0]) {

			kelime[i] = rune(degisen[0])

		} else {
			kelime[i] = rune(düzeltme[i])
		}

	}
	return string(kelime)

}

func main() {
	if len(os.Args) != 4 { // fazla argüman girdiğin zaman boş dönmesini istediğimiz durum
		fmt.Println()
		return
	}

	input := os.Args[1] // burada os ile aldığımız argümanları belirliyoruz  fonksiyonumuzda argümann  olarka verilenlerle eşleştiriryoruz
	findChar := os.Args[2]
	replaceChar := os.Args[3]

	result := ReplaceChars(input, findChar, replaceChar) // fonksiyonumuzu main içinde çalıştırıp ekrana bastırlmasını tamamladık .
	fmt.Println(result)
}
