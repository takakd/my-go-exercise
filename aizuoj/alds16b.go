package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds16b struct {
	scanner  *bufio.Scanner
	f        func(a ...interface{}) (n int, err error)
	A        []int
	sentinel int
}

//func main() {
//	a := NewAlds16b()
//	a.main()
//}

func NewAlds16b() *Alds16b {
	a := &Alds16b{}
	return a
}

func (a *Alds16b) main() {
	scanner := bufio.NewScanner(os.Stdin)
	a.handle(scanner, fmt.Print)
}

func (a *Alds16b) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	a.scanner = scanner
	a.f = f

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	a.A = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a.A[i], _ = strconv.Atoi(scanner.Text())
	}

	a.solve()

	pr := func(i int, sp string) {
		if i == a.sentinel {
			f(fmt.Sprintf("%s[%d]", sp, a.A[i]))
		} else {
			f(fmt.Sprintf("%s%d", sp, a.A[i]))
		}
	}
	i := 0
	pr(0, "")
	i++
	for i < n {
		pr(i, " ")
		i++
	}
	f("\n")
}

func (a *Alds16b) solve() {
	i, j, l := 0, 0, len(a.A)
	for j < l {
		if a.A[j] <= a.A[l-1] {
			a.A[i], a.A[j] = a.A[j], a.A[i]
			i++
		}
		j++
	}
	a.sentinel = i - 1
}
