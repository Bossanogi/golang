package arrays

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
