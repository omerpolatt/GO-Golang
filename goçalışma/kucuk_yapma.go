package main

import "fmt"

func ToLower(s string) string {

	dizi := []rune(s)
	boyut := len(dizi)
	str := ""

	for i := 0; i < boyut; i++ {
		if (dizi[i] >= 'a' && dizi[i] <= 'z') || (dizi[i] >= 'A' && dizi[i] <= 'Z') {
			if dizi[i] >= 'a' && dizi[i] <= 'z' {
				dizi[i] = dizi[i]

			} else {
				dizi[i] = dizi[i] + 32
			}
		}
		str += string(dizi[i])
	}
	return str

}
func main() {
	fmt.Println(ToLower("OMER FARUK POLAT"))
}
