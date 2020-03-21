package main

import (
	"bufio"
	"bytes"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestAlds15c(t *testing.T) {

	// input
	s := `1`
	r := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(r)

	// handle
	res := make([][]float64, 0, 0)

	f := func(a ...interface{}) (n int, err error) {
		ss := strings.Split(a[0].(string), " ")

		elm := []float64{0.0, 0.0}
		elm[0], _ = strconv.ParseFloat(ss[0], 64)
		elm[1], _ = strconv.ParseFloat(ss[0], 64)
		res = append(res, elm)

		return 0, nil
	}

	a := NewAlds15c()
	a.handle(scanner, f)

	// check
	ans := [][]float64{{0, 0}, {33.333333, 33.333333}, {50, 50}, {66.666667, 66.666667}, {100, 100}}
	for i, v := range res {
		for j := 0; j < 2; j++ {
			if math.Abs(v[j]-ans[i][j]) > 0.1 {
				t.Errorf("result was incorrect, got: %f, want: %f.", v[j], ans[i][j])
			}
		}
	}
}
