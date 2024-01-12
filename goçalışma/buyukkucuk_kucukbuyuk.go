package main

/*

package kata

import "unicode"

func ToAlternatingCase(str string) string {
  result := []rune{}
  for _, ch:=range str{
    if unicode.IsUpper(ch){
      result = append(result, unicode.ToLower(ch))
    } else if unicode.IsLower(ch){
       result = append(result, unicode.ToUpper(ch))
    } else {
      result = append(result, ch)
    }
  }

  return string(result)
}

*/

func ToAlternatingCase(str string) string {

	output := ""

	for _, harfler := range str {
		if (harfler >= 'a' && harfler <= 'z') || (harfler >= 'A' && harfler <= 'Z') {
			if harfler >= 'a' && harfler <= 'z' {
				output += string(harfler - 32)
			} else if harfler >= 'A' && harfler <= 'Z' {
				output += string(harfler + 32)
			}
		} else {
			output += string(harfler)
		}

	}
	return output
}
