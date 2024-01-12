package main

import (
	"fmt"
)

func Rot14(s string) string {
	dizi := []rune(s)
	str := ""

	for i := 0; i < len(dizi); i++ {
		if dizi[i] >= 'a' && dizi[i] <= 'z' {
			dizi[i] = ((dizi[i]-'a'+14)%26 + 'a') // verilen sayısın alfabedeki yerini bulup 14 ekle sonra mod 26 ile lafabedeki yerinin değerini belirle bu değeri a değereinin üzerine ekle
		} else if dizi[i] >= 'A' && dizi[i] <= 'Z' {
			dizi[i] = ((dizi[i]-'A'+14)%26 + 'A')
		}
		str += string(dizi[i])
	}

	return str
}

func main() {
	result := Rot14("Hello! How are You?")
	fmt.Println(result)
}
