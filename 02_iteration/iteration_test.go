package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("blah ", 3)
	fmt.Println(result)
	// Output: blah blah blah
}
