package iteration

func Repeat(term string) string {
	tmp := ""
	for i := 0; i < 5; i++ {
		tmp += term
	}

	return tmp
}
