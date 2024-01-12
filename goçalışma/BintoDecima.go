package main

import (
	"fmt"
	"math"
)

func binaryToDecimal(binary string) int {
	result := 0
	length := len(binary)

	for i := 0; i < length; i++ {
		// Binary'de en sağdan başlayarak işlem yapılır
		// Her bir bit, 2'nin üssü ile çarpılarak toplama eklenir
		bit := int(binary[length-i-1] - '0')
		result += bit * int(math.Pow(2, float64(i)))
	}

	return result
}

func main() {
	binaryNumber := "101010"
	decimalNumber := binaryToDecimal(binaryNumber)
	fmt.Println(decimalNumber)
}
