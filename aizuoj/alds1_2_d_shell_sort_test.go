package main

import (
	"testing"
)

func TestAlds12dShellSort(t *testing.T) {
	nums := []int{1, 10, 100, 11, 9, 3, 4, 2, 5, 6, 8, 7}
	n := len(nums)

	a := Alds12d{}
	a.ShellSort(nums, n)

	ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 100}
	for i, v := range nums {
		if v != ans[i] {
			t.Errorf("elemnt was incorrect, got: %d, want: %d.", v, ans[i])
		}
	}
}
