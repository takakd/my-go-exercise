package main

import "fmt"

type BubbleSort struct {
}

func (b *BubbleSort) main() {
	slice := []int{1, 4, 5, 11, 73, 3, 2, 5, 6}

	b.solve(slice)

	fmt.Println(slice)
}

func (b *BubbleSort) solve(nums []int) {
	l := len(nums)
	i := 0
	swapped := true

	for i < l-1 && swapped {

		swapped = false

		for j := l - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				swapped = true
			}
		}

		i++
	}
}
