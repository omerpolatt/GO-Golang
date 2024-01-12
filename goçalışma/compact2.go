package main

import "fmt"

func Compact(ptr *[]string) int {
	// İlk olarak, işlev bir dilim adresini alır ve bu dilimi "slice" adlı bir değişkene atar.
	slice := *ptr

	// Dilimin uzunluğunu hesaplar.
	length := len(slice)

	// Sıkıştırılmış (zero değerlere sahip öğelerin kaldırıldığı) bir dilim oluşturur.
	// Başlangıçta bu sıkıştırılmış dilim boş olacaktır.
	//son_hal := make([]string, 0)
	son_hal := []string{}

	// Döngü ile orijinal dilimdeki öğeleri tarar.
	for i := 0; i < length; i++ {
		// Eğer öğe boş bir dize değilse (sıfır olmayan bir değerse),
		// o öğeyi sıkıştırılmış dilime ekler.
		if slice[i] != "" {
			son_hal = append(son_hal, slice[i])
		}
	}

	// Sıkıştırılmış dilimi orijinal dilimin yerine atanır.
	*ptr = son_hal

	// Son olarak, sıkıştırılmış dilimin uzunluğu (sıfır olmayan öğelerin sayısı) döndürülür.
	return len(son_hal)
}

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
