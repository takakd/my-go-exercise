package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds14c struct {
	// hash array
	data []string
	// hash length
	length int
}

//func main() {
//	a := NewAlds14c()
//	a.main()
//}

func NewAlds14c() *Alds14c {
	a := &Alds14c{}
	a.length = 1<<20 - 1
	//a.length = 16777216
	a.data = make([]string, a.length)
	return a
}

func (a *Alds14c) main() {
	scanner := bufio.NewScanner(os.Stdin)
	res := a.solve(scanner)
	for _, v := range res {
		fmt.Println(v)
	}
}

func (a *Alds14c) getKey(v string) int {
	getValue := func(c byte) int {
		key := 0
		switch c {
		case 'A':
			key = 1
		case 'C':
			key = 2
		case 'G':
			key = 3
		case 'T':
			key = 4
		}
		return key
	}
	sum := 0
	p := 1
	for _, c := range []byte(v) {
		sum += p * getValue(c)
		p *= 5
	}
	return sum
}

// open address hash functions
func (a *Alds14c) h1(v int) int {
	return v % a.length
}
func (a *Alds14c) h2(v int) int {
	return 1 + (v % (a.length - 1))
}

func (a *Alds14c) insert(v string) {
	key := a.getKey(v)
	i := 0
	for {
		h := (a.h1(key) + i*a.h2(key)) % a.length
		if a.data[h] == v {
			// already inserted
			break
		} else if a.data[h] == "" {
			// insert because new value
			a.data[h] = v
			break
		}
		i++
	}
}

func (a *Alds14c) find(v string) bool {
	key := a.getKey(v)
	i := 0
	for {
		h := (a.h1(key) + i*a.h2(key)) % a.length
		if a.data[h] == v {
			return true
		}
		if a.data[h] == "" {
			return false
		}
		i++
	}
}

func (a *Alds14c) solve(scanner *bufio.Scanner) []string {

	res := make([]string, 0)

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	n *= 2

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := scanner.Text()

		scanner.Scan()
		value := scanner.Text()

		if cmd == "insert" {
			a.insert(value)
		} else if cmd == "find" {
			if a.find(value) {
				res = append(res, "yes")
			} else {
				res = append(res, "no")
			}
		}
	}

	return res
}
