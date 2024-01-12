package main

import (
	"fmt"
	"os"
)

func replaceChars(input, findChar, replaceChar string) string {
	kelime := make([]byte, len(input)) // replaced adlı bir dizi oluşturduk girilen girdi uzunluğunda .
	for i := 0; i < len(input); i++ {
		if string(input[i]) == findChar { // i deki karakter istenilen karakterle aynı ise
			kelime[i] = replaceChar[0]
		} else { // findchar ile erişmiyorsa
			kelime[i] = input[i]
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

	result := replaceChars(input, findChar, replaceChar) // fonksiyonumuzu main içinde çalıştırıp ekrana bastırlmasını tamamladık .
	fmt.Println(result)
}
