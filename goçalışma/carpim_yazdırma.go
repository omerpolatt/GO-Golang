/*package kata

import (
  "strconv"
)

func Derive(coefficient, exponent int) string {
  return strconv.Itoa(coefficient * exponent) + "x^" + strconv.Itoa(exponent - 1)
}
*/

package kata

import "strconv"

func Derive(coefficient, exponent int) string {

	if exponent != 1 && coefficient != 0 && exponent != 0 {
		carpim := coefficient * exponent
		carp := exponent - 1

		carpimStr := strconv.Itoa(carpim)
		carpStr := strconv.Itoa(carp)

		return carpimStr + "x^" + carpStr
	}

	return ""

}
