package main

import "fmt"

func main() {
	//returnArr := []int{} // Boş bir slice tanımlandı
	//var returnArr []int
	returnArr := []int
	// Nil slice'e eleman eklemek için append kullanılabilir
	returnArr = append(returnArr, 42)

	fmt.Println(returnArr) // [42]
}
