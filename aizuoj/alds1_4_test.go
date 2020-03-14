package main

import (
	"testing"
	"bufio"
	"bytes"
)

func Test14d(t *testing.T) {

	// input
	s := `5
1 2 3 4 5
3
3 4 1`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := Alds14d{}
	result := a.solve(scanner)

	// check
	ans := 3
	if result != ans {
		t.Errorf("result was incorrect, got: %d, want: %d.", result, ans)
	}
}
