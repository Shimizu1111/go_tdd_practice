package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum1(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := Sum(numbers)
	want := 6
	fmt.Println("Sum = ", Sum(numbers))
	if got != want {
		t.Errorf("want: %d but got: %d givenArray: %v", want, got, numbers)
	}
}

func TestSum2(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		if got != want {
			t.Errorf("want: %d but got: %d givenArray: %v", want, got, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		want := Sum(numbers)
		get := 6
		fmt.Println("SUM = ", Sum([]int{1, 2, 3}))
		if want != get {
			t.Errorf("want: %d get: %d", want, get)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %d but got: %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want: %d but got: %d", want, got)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
