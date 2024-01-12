// burada dizi içinde araç hızlanmalaırnı veriyor ve bu hızlanmalara göre araçın saniye saniye hızlarını bulup en büyük hızını da yazdırıyoruz

package main

import "fmt"

func Gps(s int, x []float64) int {

	if len(x) < 0 {
		return 0
	}

	max_hiz := 0.0

	for i := 1; i < len(x); i++ {

		fark := x[i] - x[i-1]

		speed := (3600 * fark) / float64(s)

		if speed > max_hiz {
			max_hiz = speed
		}

	}

	return int(max_hiz)
}

func main() {

	hiz := []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}
	s := 12

	fmt.Println(Gps(s, hiz))
}
