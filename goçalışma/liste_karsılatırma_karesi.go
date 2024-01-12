package main

import (
	"fmt"
	"sort"
)

func comp(a, b []int) bool {
	// Eğer a veya b nil ise, problem anlamsızdır, bu yüzden false döndür.
	if a == nil || b == nil {
		return false
	}

	// a ve b dizilerini sırala
	sort.Ints(a)
	sort.Ints(b)

	// b dizisinin elemanları, a dizisinin elemanlarının karelerine eşit mi kontrol et
	for i := range a {
		if b[i] != a[i]*a[i] {
			return false
		}
	}

	return true
}

func main() {
	// Örnek kullanım
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

	result := comp(a, b)
	fmt.Println(result) // Beklenen çıktı: true
}
 11  19   19 19  121     144 144 161
 121 361 361 361 14641 20736 20376  25921