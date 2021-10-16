package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 1)
	expected := 2
	if sum != expected {
		t.Errorf("expected: '%d' but got '%d'", sum, expected)
	}
}
