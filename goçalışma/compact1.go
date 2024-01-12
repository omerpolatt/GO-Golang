package main

import "fmt"

// Compact fonksiyonu, bir dilimdeki boş olmayan elemanları sıkıştırır.
// Boş olmayan elemanlar dizinin başına taşınır ve dizinin sonunda boş olmayan elemanların yeni sayısı döndürülür.
func Compact(ptr *[]string) int {
	count := 0
	for _, e := range *ptr {
		if e != "" {
			// Boş olmayan elemanları dizinin başına taşı
			(*ptr)[count] = e
			count++
		}
	}

	// Diziyi sıkıştırılmış haliyle güncelle
	*ptr = (*ptr)[0:count]
	return count
}

const N = 6

func main() {
	// N uzunluğunda bir string dilimi oluştur
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	// Oluşturulan dilimi yazdır
	fmt.Println("Original slice:")
	for _, v := range a {
		fmt.Println(v)
	}

	// Compact fonksiyonunu kullanarak dilimi sıkıştır
	fmt.Println("Size after compacting:", Compact(&a))

	// Sıkıştırılmış dilimi yazdır
	fmt.Println("Compact slice:")
	for _, v := range a {
		fmt.Println(v)
	}
}
