package main

import (
	"fmt"
	"strings"
)

type Alds11a struct {
}

func (a *Alds11a) scanInput() (n int, slice []int) {
	fmt.Scanf("%d", &n)
	slice = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &slice[i])
	}
	return
}

func (a *Alds11a) printSlice(slice []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), " "), "[]"))
}

func (a *Alds11a) main() {
	_, slice := a.scanInput()

	a.printSlice(slice)

	a.solve(slice)
}

func (a *Alds11a) solve(nums []int) {
	l := len(nums)
	for i := 1; i < l; i++ {
		v := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v

		a.printSlice(nums)
	}
}
