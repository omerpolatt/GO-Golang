// sesli harfler aeou olunca sayısnı yaazdırmak için tasarlana fonkumuz 


package main


/*

func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}

*/

func GetCount(str string) (count int) {
	sesli := "aeiou"
	
	for _, i := range str {
	  ses := false
	  for _, sesliHarf := range sesli {
		if i == sesliHarf {
		  ses = true
		  break
		}
	  }
	  if ses == true {
		count++
	  }
	}
  
	return count
  }