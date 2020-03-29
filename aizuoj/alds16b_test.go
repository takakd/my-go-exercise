package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestAlds16b(t *testing.T) {

	// input
	s := `12
13 19 9 5 12 8 7 4 21 2 6 11`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve
	res := ""
	f := func(a ...interface{}) (n int, err error) {
		res += a[0].(string)
		return 0, nil
	}
	a := NewAlds16b()
	a.handle(scanner, f)

	// check
	ans := `9 5 8 7 4 2 6 [11] 21 13 19 12
`
	if res != ans {
		t.Errorf("result was incorrect, got: %s, want: %s.", res, ans)
	}
}
