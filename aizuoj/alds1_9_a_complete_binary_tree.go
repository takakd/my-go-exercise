package main

import (
	"fmt"
)

type Alds19a struct {
}

func (a *Alds19a) scanInput() (n int, nums []int) {
	fmt.Scanf("%d", &n)
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	return
}

func (a *Alds19a) main() {
	_, nums := a.scanInput()
	length := len(nums)

	for i, v := range nums {
		node := i + 1

		parent := node / 2
		left := 2 * node
		right := 2*node + 1

		fmt.Printf("node %d: key = %d,", node, v)
		if 0 < parent && parent <= length {
			fmt.Printf(" parent key = %d,", nums[parent-1])
		}
		if left <= length {
			fmt.Printf(" left key = %d,", nums[left-1])
		}
		if right <= length {
			fmt.Printf(" right key = %d,", nums[right-1])
		}
		fmt.Println(" ")
	}
}
