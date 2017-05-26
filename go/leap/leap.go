package leap

const testVersion = 3

func IsLeapYear(n int) bool {
	return (((n%4 == 0) && (n%100 != 0)) || ((n%400 == 0) && (n%100 == 0)))
}
