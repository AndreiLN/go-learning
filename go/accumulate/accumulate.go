package accumulate

const testVersion = 1

func Accumulate(m []string, f func(string) string) []string {
	for k, v := range m {
		m[k] = f(v)
	}
	return m
}
