package main

import (
	"fmt"
	"sort"
)

func TwoSort(arr []string) string {

	// String slice'ı sırala
	sort.Strings(arr)

	var output string

	ilkKelime := arr[0]

	for _, harf := range ilkKelime {
		output += string(harf) + "***"
	}

	// Son "***" ekisini kaldır
	output = output[:len(output)-3]

	return output

}

func main() {
	arr := []string{"omer", "polat", "apple", "banana"}
	result := TwoSort(arr)
	fmt.Println(result)
}
