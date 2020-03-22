package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SENTINEL = 1000000000
)

type Alds15b struct {
	scanner      *bufio.Scanner
	f            func(a ...interface{}) (n int, err error)
	S            []int
	compareCount int
	n            int
	L            []int
	R            []int
}

//func main() {
//	a := NewAlds15b()
//	a.main()
//}

func NewAlds15b() *Alds15b {
	a := &Alds15b{}
	return a
}

func (a *Alds15b) main() {
	scanner := bufio.NewScanner(os.Stdin)
	a.handle(scanner, fmt.Print)
}

func (a *Alds15b) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	a.scanner = scanner
	a.f = f
	a.compareCount = 0

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a.n, _ = strconv.Atoi(scanner.Text())
	a.S = make([]int, a.n)
	for i := 0; i < a.n; i++ {
		scanner.Scan()
		a.S[i], _ = strconv.Atoi(scanner.Text())
	}
	// +1で配列の半分の数を確保、更に+1でSENTINEL分を確保
	a.L = make([]int, (a.n/2)+1+1)
	a.R = make([]int, (a.n/2)+1+1)

	a.mergeSort(0, a.n)

	// @note: この方法だとTLE
	//i := 0
	//f(strconv.Itoa(a.S[i]))
	//i++
	//for i < len(a.S) {
	//	f(" " + strconv.Itoa(a.S[i]))
	//	i++
	//}
	//f("\n")
	//f(strconv.Itoa(a.compareCount))
	//f("\n")

	// これだとOK
	res := strings.TrimRight(fmt.Sprintf("%+v", a.S)[1:], "]") + "\n"
	res += strconv.Itoa(a.compareCount) + "\n"
	f(res)
}

// mid, rightはend
func (a *Alds15b) mergeSort(left, right int) {
	//0 1
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2

	//fmt.Printf("left:%d, mid:%d, right:%d\n", left, mid, right)

	a.mergeSort(left, mid)
	a.mergeSort(mid, right)
	a.merge(left, mid, right)
}

func (a *Alds15b) merge(left, mid, right int) {
	// これだと要素数がスライスと同じになってしまうためSENTINELを置けない
	//a.L = a.S[left:mid]
	//a.R = a.S[mid:right]

	// midはleftのendなので、rightのbegin
	copy(a.L, a.S[left:mid])
	copy(a.R, a.S[mid:right])

	a.L[mid-left] = SENTINEL
	a.R[right-mid] = SENTINEL

	j, k := 0, 0
	for i := left; i < right; i++ {
		a.compareCount++
		// SENTINELをおいて
		// ここで、Lの末尾のときRのSENTINELと比較されるようにするところが肝
		if a.L[j] < a.R[k] {
			a.S[i] = a.L[j]
			j++
		} else {
			a.S[i] = a.R[k]
			k++
		}
	}
}
