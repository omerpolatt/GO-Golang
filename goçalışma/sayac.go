package main

import "fmt"

func CountSheeps(numbers []bool) int {
	cnt := 0
	for _, b := range numbers {
		if b {
			cnt++
		}
	}
	return cnt
}

func main() {
	// Örnek kullanım
	sheepArray := []bool{
		false, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
		// nil değeri eklenmemiş durumda
	}

	result := CountSheeps(sheepArray)
	fmt.Println(result)
}
