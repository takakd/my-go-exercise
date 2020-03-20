package main

import (
	"bufio"
	"bytes"
	"testing"
)

func Test14c(t *testing.T) {

	// input
	s := `15
insert AAA
insert AAC
insert AGA
insert AGG
insert TTT
find AAA
find CCC
find CCC
insert CCC
find CCC
find CC
insert T
find TTT
find T
find A
`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	a := NewAlds14c()
	tmp := a.solve(scanner)

	result := ""
	for _, v := range tmp {
		result += v + "\n"
	}

	// check
	ans :=
		`yes
no
no
yes
no
yes
yes
no
`
	if result != ans {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, ans)
	}
}
