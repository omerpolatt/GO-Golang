package main

import (
	"fmt"
	"os"
)

func Greet(name string) string {
	var output string

	if name == "Johnny" {
		output = "Hello, my love!"

	} else {
		output = "Hello, " + name + "!"
	}
	return output
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		cikti := Greet(args[0])
		fmt.Println(cikti)
	}
}
