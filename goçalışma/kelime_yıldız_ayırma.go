package main

import (
	"fmt"
	"sort"
)

func TwoSort(arr []string) string {
	sort.Strings(arr) // String slice'ı sırala

	var output string

	for i, word := range arr {
		if i > 0 {
			output += "***"
		}
		output += word
	}

	return output
}

func main() {
	arr := []string{"omer", "polat", "apple", "banana"}
	result := TwoSort(arr)
	fmt.Println(result)
}
