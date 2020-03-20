package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds15a struct {
	a []int
	n int
	m []int
	q int
}

//func main() {
//	a := NewAlds15a()
//	a.main()
//}

func NewAlds15a() *Alds15a {
	a := &Alds15a{}
	return a
}

func (a *Alds15a) main() {
	scanner := bufio.NewScanner(os.Stdin)
	a.handle(scanner, fmt.Println)
}

func (a *Alds15a) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a.n, _ = strconv.Atoi(scanner.Text())

	a.a = make([]int, a.n)
	for i := 0; i < a.n; i++ {
		scanner.Scan()
		a.a[i], _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	a.q, _ = strconv.Atoi(scanner.Text())

	a.m = make([]int, a.q)
	for i := 0; i < a.q; i++ {
		scanner.Scan()
		a.m[i], _ = strconv.Atoi(scanner.Text())
	}

	for _, v := range a.m {
		if a.solve(0, v) {
			f("yes")
		} else {
			f("no")
		}
	}
}

func (a *Alds15a) solve(i, m int) bool {
	if m == 0 {
		return true
	}

	if a.n <= i {
		return false
	}
	return a.solve(i+1, m) || a.solve(i+1, m-a.a[i])
}
