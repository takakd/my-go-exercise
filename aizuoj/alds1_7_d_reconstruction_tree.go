package main

import (
	"fmt"
	"strings"
)

type Alds17d struct {
}

func (a *Alds17d) scanInput() (n int, pre []int, in []int) {
	fmt.Scanf("%d", &n)
	pre = make([]int, n)
	in = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &pre[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &in[i])
	}
	return
}

func (a *Alds17d) printSlice(slice []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), " "), "[]"))
}

func (a *Alds17d) main() {
	n, pre, in := a.scanInput()
	post := make([]int, 0, n)
	parent := 0

	a.walk(0, len(pre)-1, &parent, pre, in, &post)

	a.printSlice(post)
}

func (a *Alds17d) walk(left, right int, parent *int, pre, in []int, post *[]int) {
	if left > right {
		return
	}

	dlm := pre[*parent]
	*parent++

	var idx int
	for i, v := range in {
		if v == dlm {
			idx = i
			break
		}
	}

	a.walk(left, idx-1, parent, pre, in, post)
	a.walk(idx+1, right, parent, pre, in, post)

	*post = append(*post, dlm)
}
