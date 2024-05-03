package iteration

func Repeat(in string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += in
	}
	return
}
