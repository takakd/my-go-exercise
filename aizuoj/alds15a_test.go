package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestAlds15a(t *testing.T) {

	// input
	s := `5
1 5 7 10 21
4
2 4 17 8`

	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// solve

	res := ""
	f := func(a ...interface{}) (n int, err error) {
		res += a[0].(string) + "\n"
		return 0, nil
	}

	a := NewAlds15a()
	a.handle(scanner, f)

	// check
	ans :=
		`no
no
yes
yes
`
	if res != ans {
		t.Errorf("result was incorrect, got: %s, want: %s.", res, ans)
	}
}
