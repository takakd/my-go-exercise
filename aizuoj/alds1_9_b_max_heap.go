package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds19b struct {
}

func (a *Alds19b) scanInput() (n int, nums []int) {
	//// NG
	//fmt.Scanf("%d", &n)
	//nums = make([]int, n)
	//for i := 0; i < n; i++ {
	//    fmt.Scanf("%d", &nums[i])
	//}

	// OK
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	return
}

func (a *Alds19b) maxHeapify(nums *[]int, length int, previousMax int) {
	parent := previousMax - 1
	left := 2*previousMax - 1
	right := 2*previousMax + 1 - 1

	if left > length-1 {
		return
	}

	largest := parent
	if left < length && (*nums)[parent] < (*nums)[left] {
		largest = left
	}
	if right < length && (*nums)[largest] < (*nums)[right] {
		largest = right
	}

	if largest != parent {
		(*nums)[largest], (*nums)[parent] = (*nums)[parent], (*nums)[largest]
		a.maxHeapify(nums, length, largest+1)
	}
}

func (a *Alds19b) main() {
	length, nums := a.scanInput()

	for i := length / 2; i >= 1; i-- {
		a.maxHeapify(&nums, length, i)
	}
	for _, v := range nums {
		fmt.Printf(" %d", v)
	}
	fmt.Println("")
}
