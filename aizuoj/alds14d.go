package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds14d struct {
	packet []int
}

//func main() {
//	a := &Alds14d{}
//	a.main()
//}

func NewAlds14d() *Alds14d {
	a := &Alds14d{}
	return a
}

func (a *Alds14d) main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(a.solve(scanner))
}

// can load in k tracks of capacity P that can carry n packet
func (a *Alds14d) calcMaxPacketNumber(P, k, n int) int {
	// packet num
	i := 0

	// 1. トラックごとに
	// per track
	for j := 0; j < k; j++ {

		// total of weight
		s := 0

		// 2. 詰めるだけpacketを積む
		for s+a.packet[i] <= P {
			s += a.packet[i]

			i++

			// 3. パケットの最大数までたどりついたらおわり
			// end of packet num
			if i == n {
				// can load n packet
				return n
			}
		}
	}

	// 4. 最大までたどり着かなかったら、途中までの個数が積むことができた個数
	return i
}

func (a *Alds14d) solve(scanner *bufio.Scanner) int {

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	a.packet = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a.packet[i], _ = strconv.Atoi(scanner.Text())
	}

	// over limit
	//capacity := 0
	//for {
	//	capacity++
	//
	//	num, valid := a.calcRequiredNumber(capacity, n)
	//	if ! valid {
	//		continue
	//	}
	//	if num <= k {
	//		break
	//	}
	//}

	left := 1
	right := 100000 * 100000
	mid := (left + right) / 2
	for right-left > 0 {
		num := a.calcMaxPacketNumber(mid, k, n)

		if n == num {
			// 必要な個数を乗せることができるので、Pを下げて検索
			// i == nの判定を行っているので、numがnを超えることはない
			right = mid
		} else {
			// Pが不足するので、より大きい値で検索
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return mid
}
