package main

import (
	"testing"
	"bufio"
	"bytes"
)

func Test14d(t *testing.T) {

	// input
	s := `5 3
8
1
7
3
9`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := NewAlds14d()
	result := a.solve(scanner)

	// check
	ans := 10
	if result != ans {
		t.Errorf("result was incorrect, got: %d, want: %d.", result, ans)
	}
}
