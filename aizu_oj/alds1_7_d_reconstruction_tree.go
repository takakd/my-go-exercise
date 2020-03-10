package main

import (
	"fmt"
	"strings"
)

func scanInput() (n int, pre []int, in []int) {
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

func printSlice(slice []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), " "), "[]"))
}

func main() {
	n, pre, in := scanInput()
	post := make([]int, 0, n)
	parent := 0

	walk(0, len(pre)-1, &parent, pre, in, &post)

	printSlice(post)
}

func walk(left, right int, parent *int, pre, in []int, post *[]int) {
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

	walk(left, idx-1, parent, pre, in, post)
	walk(idx+1, right, parent, pre, in, post)

	*post = append(*post, dlm)
}

