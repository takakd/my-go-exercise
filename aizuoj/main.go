package main

import (
	"fmt"
	"strings"
)

type Main struct {
}

func (m *Main) main() {
	var n int
	fmt.Scanf("%d", &n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &values[i])
	}
	fmt.Println(n, values)
}

func (m *Main) arrayToString(a []int, delim string) string {
	//return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
