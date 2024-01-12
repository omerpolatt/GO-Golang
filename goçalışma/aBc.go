package main

import "github.com/01-edu/z01"

func main() {

	sayac := 0

	for i := 'a'; i <= 'z'; i++ {
		if sayac%2 == 0 {
			z01.PrintRune(i)
			sayac++

		} else {
			z01.PrintRune(i - 32)
			sayac++
		}
	}
}
