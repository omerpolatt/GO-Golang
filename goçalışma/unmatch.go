package main

import "fmt"

func Unmatch(a []int) []int {
	sayac := 0
	sonuc := []int{}

	for _, i := range a {

		for _, ş := range a {
			if i == ş {
				sayac++
			}
		}

		if sayac == 1 || sayac%2 == 1 {
			sonuc = append(sonuc, i)

		}

	}
	return sonuc

}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 1}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
