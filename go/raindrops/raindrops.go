package raindrops

import "strconv"

const testVersion = 3

func Convert(n int) string {
	var result string
	result = concat(result, n, 3, "Pling")
	result = concat(result, n, 5, "Plang")
	result = concat(result, n, 7, "Plong")
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}
func concat(s string, number int, divisor int, word string) string {
	if number%divisor == 0 {
		s = s + word
	}
	return s
}
