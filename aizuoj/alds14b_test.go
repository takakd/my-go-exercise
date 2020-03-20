package main

import (
	"bufio"
	"bytes"
	"testing"
)

func Test14b(t *testing.T) {

	// input
	s := `15
0 0 1 1 2 2 3 3 3 5 6 6 8 9 9
10
8 4 6 5 0 2 1 7 9 3
`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := Alds14b{}
	result := a.solve(scanner)

	// check
	ans := 8
	if result != ans {
		t.Errorf("result was incorrect, got: %d, want: %d.", result, ans)
	}
}
