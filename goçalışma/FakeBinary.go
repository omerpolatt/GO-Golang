package main

import (
	"fmt"
	"strconv"
)

func FakeBin(x string) string {

	output := ""

	for _, i := range x {
		sayi, err := strconv.Atoi(string(i))

		if err != nil {
			fmt.Println(err)
		} else {

			if sayi >= 5 {
				output += "1"
			} else {
				output += "0"
			}
		}
	}
	return output
}
