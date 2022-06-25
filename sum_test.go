package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("合計値の値が正しい", func(t *testing.T) {
		got := Sum(1, 2)
		want := 3

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}