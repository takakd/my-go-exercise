package main

import (
	"bufio"
	"bytes"
	"testing"
)

func Test14a(t *testing.T) {

	// input
	s := `5
1 2 3 4 5
3
3 4 1`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := Alds14a{}
	result := a.solve(scanner)

	// check
	ans := 3
	if result != ans {
		t.Errorf("result was incorrect, got: %d, want: %d.", result, ans)
	}
}
