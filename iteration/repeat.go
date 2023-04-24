package iteration

const repeatCount = 5

func Repeat(c string, n int) string {
	var s string

	if n == 0 {
		n = repeatCount
	}

	for l := 0; l < n; l++ {
		s += c
	}
	return s
}
