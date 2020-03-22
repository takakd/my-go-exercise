package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestAlds15b(t *testing.T) {

	// input
	s := `10
8 5 9 2 6 3 7 1 10 4`

	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	res := ""
	f := func(a ...interface{}) (n int, err error) {
		res += a[0].(string)
		return 0, nil
	}
	a := NewAlds15b()
	a.handle(scanner, f)

	// check
	ans := `1 2 3 4 5 6 7 8 9 10
34
`
	if res != ans {
		t.Errorf("result was incorrect, got: %s, want: %s.", res, ans)
	}
}
