// @ref https://blog.alexellis.io/golang-writing-unit-tests/
package main

import "testing"

// Test of Sum
func TestSum(t *testing.T) {
	tables := []struct {
		x, y, n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
		{5, 2, 9},
	}

	for _, table := range tables {
		r := Sum(table.x, table.y)
		if r != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, r, table.n)

		}
	}
}
