package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func scanInput() (n int, nums []int) {
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

func maxHeapify(nums *[]int, length int, previousMax int) {
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
		maxHeapify(nums, length, largest+1)
	}
}

func main() {
	length, nums := scanInput()

	for i := length / 2; i >= 1; i-- {
		maxHeapify(&nums, length, i)
	}
	for _, v := range nums {
		fmt.Printf(" %d", v)
	}
	fmt.Println("")
}

