package iteration

func Repeat(term string) (repeated string) {
	repeated = ""
	for i := 0; i < 5; i++ {
		repeated += term
	}

	return
}
