package main

/*
 result := make([]int, 0, limit/integer)
	for i := integer; i <= limit; i += integer {
		result = append(result, i)
	}

	return result
*/

func FindMultiples(integer, limit int) []int {

	dizi := []int{}

	for i := integer; i <= limit; i++ {
		if i%integer == 0 && i > 0 {
			dizi = append(dizi, i)
		}
	}
	return dizi
}
