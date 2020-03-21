package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Alds15c struct {
	f func(a ...interface{}) (n int, err error)
}

//func main() {
//	a := NewAlds15c()
//	a.main()
//}

func NewAlds15c() *Alds15c {
	a := &Alds15c{}
	return a
}

func (a *Alds15c) main() {
	scanner := bufio.NewScanner(os.Stdin)
	a.handle(scanner, fmt.Println)
}

type Point struct {
	X float64
	Y float64
}

// コッホ曲線を出力
// 内分点の結果をa.fへ渡す
func (a *Alds15c) koch(d int, p1 *Point, p2 *Point) {
	if d == 0 {
		return
	}

	s := &Point{X: ((2 * p1.X) / 3) + (p2.X / 3), Y: ((2 * p1.Y) / 3) + (p2.Y / 3)}
	t := &Point{X: ((p1.X) / 3) + (2 * p2.X / 3), Y: ((p1.Y) / 3) + (2 * p2.Y / 3)}

	// rotate vector(s -> u)
	rad := math.Pi / 3
	cos60 := math.Cos(rad)
	sin60 := math.Sin(rad)
	st := &Point{X: t.X - s.X, Y: t.Y - s.Y}
	u := &Point{X: (st.X*cos60 - st.Y*sin60) + s.X, Y: (st.X*sin60 + st.Y*cos60) + s.Y}

	a.koch(d-1, p1, s)
	a.f(fmt.Sprintf("%f %f", s.X, s.Y))

	a.koch(d-1, s, u)
	a.f(fmt.Sprintf("%f %f", u.X, u.Y))

	a.koch(d-1, u, t)
	a.f(fmt.Sprintf("%f %f", t.X, t.Y))

	a.koch(d-1, t, p2)
}

func (a *Alds15c) handle(scanner *bufio.Scanner, f func(a ...interface{}) (n int, err error)) {

	a.f = f

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	p1 := &Point{X: 0.0, Y: 0.0}
	p2 := &Point{X: 100.0, Y: 0.0}

	// 始点
	a.f(fmt.Sprintf("%f %f", p1.X, p1.Y))

	// コッホ曲線の内分点
	a.koch(d, p1, p2)

	// 終点
	a.f(fmt.Sprintf("%f %f", p2.X, p2.Y))
}
