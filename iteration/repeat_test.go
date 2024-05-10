package iteration

import "testing"

const times = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", times)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", times)
	}
}
