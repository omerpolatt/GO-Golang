package main

import "fmt"

func Disemvowel(comment string) string {
	output := ""
	sesliHarfler := "aeıioöuüAEIİOÖUÜ"

	for _, harf := range comment {
		sesli := false

		for _, sesliHarf := range sesliHarfler {
			if harf == sesliHarf {
				sesli = true
				break
			}
		}

		if !sesli {
			output += string(harf)
		}
	}

	return output
}

func main() {

	yazi := "omer faruk"

	fmt.Println(Disemvowel(yazi))

}
