// @ref https://blog.alexellis.io/golang-writing-unit-tests/
package main

import "fmt"

// Sum returns sum of x and b.
func Sum(x int, y int) int {
	return x + y
}

func main() {
	r := Sum(5, 5)
	fmt.Println(r)
}
