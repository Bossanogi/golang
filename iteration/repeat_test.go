package iteration

import (
	"fmt"
	"testing"
)

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

func ExampleRepeat() {
	repeating := Repeat("ab", 3)
	fmt.Println(repeating)
}

func TestRepeatC(t *testing.T) {
	repeated := RepeatC("a", times)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
