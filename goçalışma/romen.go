package main

import (
	"fmt"
	"os"
	"strconv"
)

// romanDigit represents a Roman numeral and its corresponding value.
type romanDigit struct {
	Value  int
	Symbol string
}

var romanDigits = []romanDigit{
	{1000, "M"},
	{900, "(M-C)"},
	{500, "D"},
	{400, "(D-C)"},
	{100, "C"},
	{90, "(C-X)"},
	{50, "L"},
	{40, "(L-X)"},
	{10, "X"},
	{9, "(X-I)"},
	{5, "V"},
	{4, "(V-I)"},
	{1, "I"},
}

// toRoman converts an Arabic numeral to a Roman numeral.
func toRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", fmt.Errorf("ERROR: cannot convert to roman digit")
	}

	roman := ""
	for _, digit := range romanDigits {
		for num >= digit.Value {
			roman += digit.Symbol // en büyük yakın değere geliyor 12 için 10 a geliyor x gibi sonra 12 den 10 u çıkar klaan değerleri oluşturacak max değerden yakın olacak şekilde oluştuuryor
			num -= digit.Value
		}
	}

	return roman, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <number>")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, err := toRoman(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(roman)
}
