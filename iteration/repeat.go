package iteration

func Repeat(s string) string {
	str := ""
	for l := 0; l < 5; l++ {
		str += s
	}
	return str
}
