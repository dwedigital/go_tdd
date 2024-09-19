package main

import (
	"math/rand"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		assertCorrectResult(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{}
		test := 0

		for i := 0; i < rand.Intn(100); i++ {
			numbers = append(numbers, i)
		}

		for _, num := range numbers {
			test += num
		}

		got := Sum(numbers)
		want := test

		assertCorrectResult(t, got, want, numbers)
	})
}

func assertCorrectResult(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
