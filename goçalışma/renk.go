package main

import "fmt"

func Triangle(row string) rune {
	if len(row) != 0 {
		currentRes := row

		for len(currentRes) > 1 {
			res := "" // yeni satır ekleyip her renk kombinasyonundaki sonucu buna ekliyor
			for i := 0; i < len(currentRes)-1; i++ {
				if currentRes[i] == currentRes[i+1] {
					res += string(currentRes[i])
				} else if (currentRes[i] == 'B' && currentRes[i+1] == 'G') || (currentRes[i] == 'G' && currentRes[i+1] == 'B') {
					res += "R"
				} else if (currentRes[i] == 'R' && currentRes[i+1] == 'G') || (currentRes[i] == 'G' && currentRes[i+1] == 'R') {
					res += "B"
				} else if (currentRes[i] == 'B' && currentRes[i+1] == 'R') || (currentRes[i] == 'R' && currentRes[i+1] == 'B') {
					res += "G"
				}
			}
			fmt.Println(res)
			currentRes = res
		}

		r := []rune(currentRes)
		return r[0]
	}
	return 'A'
}

func main() {
	renk := "RRGBRGBB"

	output := Triangle(renk)

	fmt.Println("Sonuç:", string(output))
}
