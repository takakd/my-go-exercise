package main

import (
	"bufio"
	"bytes"
	"testing"
)

func Test13d(t *testing.T) {

	// input
	s := `\\///\_/\/\\\\/_/\\///__\\\_\\/_\/_/\`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := Alds13d{}
	result := a.solve(scanner)

	// check
	ans :=
		`35
5 4 2 1 19 9
`
	if result != ans {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, ans)
	}
}
