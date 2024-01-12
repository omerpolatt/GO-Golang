package main

import (
	"fmt"
	"os"
	"strconv"
)

func asal_bulma(nb int) bool {
	sayac := 0

	if nb < 2 {
		return false
	}

	for i := 2; i <= nb/2; i++ {

		if nb%i == 0 {
			sayac++
		}
	}
	if sayac == 0 {
		return true
	}
	return false

}

func fark(slice []int, eleman int) bool {
	for _, s := range slice {
		if s == eleman {
			return true
		}
	}
	return false
}

func asal_alma(sayilar []int) []int {

	dizi := []int{}

	for _, i := range sayilar {

		if fark(dizi, i) == false && asal_bulma(i) == true {
			dizi = append(dizi, i)
		}
	}

	return dizi

}

func main() {

	args := os.Args[1:]

	sayilar := []int{}

	for _, arg := range args {

		sayi, err := strconv.Atoi(arg)

		if err != nil {
			fmt.Println("hata aldik", err)
			return

		}

		sayilar = append(sayilar, sayi)

	}
	asal_sayilar := asal_alma(sayilar)

	fmt.Println(asal_sayilar)

}
