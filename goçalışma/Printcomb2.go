package main

import "fmt"

func PrintComb2() {
	// İki haneli tüm rakamları döngü ile gez
	for i := '0'; i <= '9'; i++ {
		// İki haneli tüm rakamları bir başka döngü ile gez
		for j := '0'; j <= '9'; j++ {
			// İkinci haneli sayıyı temsil etmek üzere f'yi tanımla ve j'den bir sonraki rakamla başlat
			f := rune(j) + 1

			// İki haneli tüm kombinasyonları oluştur
			for k := i; k <= '9'; k++ {
				for ; f <= '9'; f++ {
					// Kombinasyonu ekrana yazdır
					fmt.Printf("%c%c %c%c", i, j, k, f)

					// Son kombinasyon değilse virgül ve boşluk ekleyerek ayrı devam et
					if i < '9' || j < '8' || k < '9' || f < '9' {
						fmt.Print(", ")
					}
				}
				// f'yi sıfırla
				f = '0'
			}
		}
	}
	// Bir alt satıra geç
	fmt.Println()
}

func main() {
	PrintComb2()
}
