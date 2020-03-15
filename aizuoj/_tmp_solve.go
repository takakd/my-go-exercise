package main

import (
	"fmt"
	"bufio"
	"os"
)

type Alds13d struct {
}

//func main() {
//	a := &Alds13d{}
//	a.main()
//}

func (a *Alds13d) main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(a.solve(scanner))
}

func (a *Alds13d) solve(scanner *bufio.Scanner) string {

	sum := 0
	s1 := make([]int, 0)
	s2 := make([]struct {
		I int
		S int
	}, 0)

	i := 0
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	for _, r := range scanner.Text() {
		c := string(r)

		if c == "\\" {
			s1 = append(s1, i)

		} else if c == "/" && len(s1) > 0 {
			j := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			s := i - j

			sum += s

			t := s
			if len(s2) > 0 {
				v := s2[len(s2)-1]
				for v.I > j {
					t += v.S
					s2 = s2[:len(s2)-1]

					if len(s2) == 0 {
						break
					}

					v = s2[len(s2)-1]
				}
			}
			s2 = append(s2, struct {
				I int
				S int
			}{j, t})
		}

		scanner.Scan()
		c = scanner.Text()
		i++
	}

	// output
	result := fmt.Sprintln(sum)
	result += fmt.Sprintf("%d", len(s2))
	for _, v := range s2 {
		result += fmt.Sprintf(" %d", v.S)
	}
	result += fmt.Sprintln()

	return result
}
