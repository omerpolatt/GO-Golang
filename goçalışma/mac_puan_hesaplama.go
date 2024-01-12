/*
package kata

func Points(games []string) int {
  points := 0
  for _, g := range games {
      if g[0]>g[2] {
        points += 3
      } else if g[0] == g[2] {
        points += 1
      }
  }
  return points
}
*/

/*

func Points(games []string) int {
	puan := 0

	for i := 0; i < len(games); i++ {
		evSahibiSkor := int(games[i][0] - '0') // İlk karakteri int'e çevir
		misafirSkor := int(games[i][2] - '0') // Üçüncü karakteri int'e çevir

		if evSahibiSkor > misafirSkor {
			puan += 3
		} else if evSahibiSkor < misafirSkor {
			puan += 0
		} else {
			puan += 1
		}
	}

	return puan
}


*/



package kata

func Points(games []string) int {
  puan := 0
  
  for i:=0 ; i<len(games);i++{
    
    if [i][0] > [i][2]{
      puan += 3
      
    }else if  [i][0] < [i][2]{
      puan += 0
    }else{
      puan +=1
    }
    return puan
  }
  return -1
}