package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("testing func Add", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})
	t.Run("testing func Sub", func(t *testing.T) {
		substruct := Sub(3, 1)
		expected := 2
		assertCorrectMessage(t, substruct, expected)
	})
}

func assertCorrectMessage(t testing.TB, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
