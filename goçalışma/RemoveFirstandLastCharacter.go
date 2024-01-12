package main

import (
	"fmt"
	"os"
)

/*
if len(word) > 2 {
		result = word[1 : len(word)-1]
	}
	return

*/

/*
return word[1:len(word)-1]
*/

func RemoveChar(word string) string {
	var output string

	if len(word) < 2 {
		output = ""
	} else {

		for i := 1; i < len(word)-1; i++ {
			output = output + string(word[i])
		}

	}
	return output

}
func main() {

	args := os.Args[1:]

	if len(args) == 1 {
		cikti := RemoveChar(args[0])
		fmt.Println(cikti)
	}

}
