package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := &Alds12d{}
	a.main()
}

type Alds12d struct {
	swapCount int
	G         []int
}

// test
func (a *Alds12d) test() {

	nums := []int{5, 1, 4, 3, 2}
	//n := len(nums)

	a.swapCount = 0
	a.ShellSort(nums, len(nums))

	fmt.Println(len(a.G))
	for _, v := range a.G {
		fmt.Print(strconv.Itoa(v) + " ")
	}
	fmt.Println("")

	fmt.Println(a.swapCount)

	for _, v := range nums {
		fmt.Println(strconv.Itoa(v) + " ")
	}
}

// get input value
func (a *Alds12d) scanInput() (n int, nums []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	return
}

// shell sort
func (a *Alds12d) ShellSort(nums []int, n int) {

	// make G
	// g(n+1) = 1 + 3 * g(n)
	a.G = []int{}
	v := 1
	for v <= n {
		a.G = append(a.G, v)
		v = 1 + 3*v
	}

	// Gの各値絵挿入ソートを実行
	// 最後は1で挿入ソートを実行するが、そのタイミングではほぼソートされていることが期待される
	for i := len(a.G) - 1; i >= 0; i-- {

		// gの間隔で、配列の後ろから挿入ソート
		g := a.G[i]

		// このループだと不要なswapが発生する
		//for r := g; r < n; r++ {
		//
		//	l := r - g
		//
		//	if nums[l] > nums[r] {
		//		nums[l], nums[r] = nums[r], nums[l]
		//		a.swapCount++
		//	}
		//
		//}

		for r := g; r < n; r++ {

			// gの間だけ開けた左側にある比較対象
			l := r - g

			// rの位置の値を退避
			v := nums[r]

			// rからgだけ左の値と比較、その間隔で戦闘まで繰り返し、左が大きければ入れ替え
			for l >= 0 && nums[l] > v {
				nums[l+g] = nums[l]
				a.swapCount++

				// g間隔ずつ移動
				l = l - g
			}

			// ループの最後の移動を打ち消し
			l = l + g
			nums[l] = v
		}
	}
}

// main
func (a *Alds12d) main() {

	//test()

	// input
	n, nums := a.scanInput()

	// solve
	a.ShellSort(nums, n)

	// output

	// - NG: OLE is raised
	//fmt.Println(len(g))
	//for i := len(g) - 1; i >= 0; i-- {
	//	fmt.Print(strconv.Itoa(g[i]) + " ")
	//}
	//fmt.Println("")
	//
	//fmt.Println(a.swapCount)
	//
	//for _, v := range nums {
	//	fmt.Println(strconv.Itoa(v) + " ")
	//}

	// OK
	fmt.Println(len(a.G))
	for i := len(a.G) - 1; i >= 0; i-- {
		fmt.Printf("%d ", a.G[i])
	}
	fmt.Println("")

	fmt.Println(a.swapCount)

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}
