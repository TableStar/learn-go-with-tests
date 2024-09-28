package main

import (
	"testing"

	"slices"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//guide used reflect.DeepEqual but go v1.21 has slice standard package which has slices.Equal function to do a simple shallow compare on slices
	//Note that this function expects the elements to be comparable. So, it can't be applied to slices with non-comparable elements like 2D slices.
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
