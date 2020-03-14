package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Alds14d struct {
}

//func main() {
//	a := &Alds14d{}
//	a.main()
//}

func (a *Alds14d) main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(a.solve(scanner))
}

func (a *Alds14d) solve(scanner *bufio.Scanner) int {

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	S := make([]int, n+1)
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

	i := 0
	for i < q {
		S[n] = T[i]

		j := 0
		for S[j] != T[i] {
			j++
		}
		if j != n {
			res++
		}

		i++
	}

	return res
}
