package main

import (
	"fmt"
	"strings"
)

type SortInsert struct {
}

func (s *SortInsert) scanInput() (n int, slice []int) {
	fmt.Scanf("%d", &n)
	slice = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &slice[i])
	}
	return
}

func (s *SortInsert) printSlice(slice []int) {
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), " "), "[]"))
}

func (s *SortInsert) main() {
	_, slice := s.scanInput()

	s.printSlice(slice)

	s.solve(slice)

	s.printSlice(slice)
}

func (s *SortInsert) solve(nums []int) {
	l := len(nums)

	i := 0
	for i < l-1 {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		i++
	}
}
