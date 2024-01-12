package main

func Solution(a, b string) string {

	var sonuc string

	if len(a) > len(b) {

		sonuc = b + a + b
	}

	if len(b) > len(a) {

		sonuc = a + b + a
	}
	return sonuc

}
