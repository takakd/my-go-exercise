package main

import (
	"fmt"
	"strings"
)

func scanInput() (n int, slice []int) {
	fmt.Scanf("%d", &n)
	slice = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &slice[i])
	}
	return
}

func printSlice(slice []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), " "), "[]"))
}

func main() {
	_, slice := scanInput()

	printSlice(slice)

	solve(slice)
}

func solve(nums []int) {
	l := len(nums)
	for i := 1; i < l; i++ {
		v := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v

		printSlice(nums)
	}
}

