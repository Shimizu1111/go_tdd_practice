package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	fmt.Println("repeat = ", Repeat("a"))
	if repeated != expected {
		t.Errorf("expected: %q but repeated %q", expected, repeated)
	}
}

func BenchmarkPepeat(t *testing.B) {
	for i := 0; i < 5; i++ {
		Repeat("a")
	}
}
