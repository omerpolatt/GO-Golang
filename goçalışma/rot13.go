package main

import (
	"fmt"
	"os"
)

func Rot13(s string) string {
	str := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			str += string('a' + ((r - 'a' + 13) % 26))
		} else if r >= 'A' && r <= 'Z' {
			str += string('A' + ((r - 'A' + 13) % 26))
		} else {
			str += string(r)
		}
	}
	return str
}
func main() {
	yazi := os.Args[1]
	cikti := Rot13(yazi)
	fmt.Println(cikti)
}
