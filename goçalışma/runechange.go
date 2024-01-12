package main

import (
	"fmt"
	"os"
)

func toWeirdCase(str string) string {
	dizi := []rune(str)
	boyut := len(dizi)
	cikti := ""
	ilk_index := 0 // İlk harfin büyük olacağı indeksi belirlemek için

	for i := 0; i < boyut; i++ {
		if dizi[i] == ' ' {
			cikti = cikti + string(dizi[i])
			ilk_index = 0 // Boşluktan sonra yeni kelimenin başındayız
		} else {
			if ilk_index%2 == 0 {
				if dizi[i] >= 'a' && dizi[i] <= 'z' {
					dizi[i] = dizi[i] - 32
				}
			} else {
				if dizi[i] >= 'A' && dizi[i] <= 'Z' {
					dizi[i] = dizi[i] + 32
				}
			}
			cikti = cikti + string(dizi[i])
			ilk_index++
		}
	}

	return cikti
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		result := toWeirdCase(args[0])
		fmt.Println(result)
	}
}
