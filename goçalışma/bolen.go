package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsDivisible(n, x, y int) bool {

	if n > 0 && x > 0 && y > 0 {

		if n%x == 0 && n%y == 0 {
			return true
		}

	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) == 3 {
		// String'leri int'e çevir
		n, errN := strconv.Atoi(args[0])
		x, errX := strconv.Atoi(args[1])
		y, errY := strconv.Atoi(args[2])

		// Hata kontrolü yap
		if errN == nil && errX == nil && errY == nil {
			output := IsDivisible(n, x, y)
			fmt.Println(output)
		} else {
			fmt.Println("Hata: Geçersiz argümanlar")
		}
	} else {
		fmt.Println("Hata: Argüman sayısı eksik veya fazla")
	}

}
