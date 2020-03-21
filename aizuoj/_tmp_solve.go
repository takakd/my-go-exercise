package main

import (
	"bufio"
	"fmt"
	"os"
)

type AldsTmp struct {
	f func(a ...interface{}) (n int, err error)
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
	fmt.Println(a.solve(scanner))
}

func (a *AldsTmp) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	a.f = f

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a.n, _ = strconv.Atoi(scanner.Text())

	// a.solve(scanner)
}

//func (a *AldsTmp) solve(scanner *bufio.Scanner) string {
//}
