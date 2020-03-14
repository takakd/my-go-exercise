package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Alds14b struct {
}

//func main() {
//	a := &Alds14b{}
//	a.main()
//}

func (a *Alds14b) main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(a.solve(scanner))
}

func (a *Alds14b) solve(scanner *bufio.Scanner) int {

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	S := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		S[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	T := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		T[i], _ = strconv.Atoi(scanner.Text())
	}

	res := 0
	l := len(S)
	for i := 0; i < q; i++ {
		//left, right := 0, l - 1
		//for left <= right {
		//	mid := (left + right) / 2
		//	if S[mid] == T[i] {
		//		res++
		//		break
		//	} else if S[mid] < T[i] {
		//		left = mid + 1
		//	} else {
		//		right = mid - 1
		//	}
		//}
		left, right := 0, l
		for left < right {
			mid := (left + right) / 2
			if S[mid] == T[i] {
				res++
				break
			} else if S[mid] < T[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return res
}
