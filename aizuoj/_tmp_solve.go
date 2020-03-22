package main

import (
	"bufio"
	"fmt"
	"os"
)

type AldsTmp struct {
	scanner *bufio.Scanner
	f       func(a ...interface{}) (n int, err error)
}

//func main() {
//	a := NewAldsTmp()
//	a.main()
//}

func NewAldsTmp() *AldsTmp {
	a := &AldsTmp{}
	return a
}

func (a *AldsTmp) main() {
	scanner := bufio.NewScanner(os.Stdin)
	a.handle(scanner, fmt.Println)
}

func (a *AldsTmp) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	a.scanner = scanner
	a.f = f

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	// n, _ = strconv.Atoi(scanner.Text())

	// a.solve()
}

//func (a *AldsTmp) solve() string {
//}
