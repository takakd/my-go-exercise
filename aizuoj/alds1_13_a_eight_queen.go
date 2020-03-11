package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Alds113a struct {
	row, col   [N]int
	dpos, dneg [2*N - 1]int
	X          [N][N]bool
}

const (
	N        = 8
	FREE     = -1
	NOT_FREE = -2
)

func (a *Alds113a) initialize() {
	for i := 0; i < N; i++ {
		a.row[i], a.col[i] = FREE, FREE
	}
	for i := 0; i < 2*N-1; i++ {
		a.dpos[i], a.dneg[i] = FREE, FREE
	}
}

func (a *Alds113a) print() {
	// 回答となる配置パターンの数だけprintがコールされる
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a.X[i][j] {
				if a.row[i] != j {
					// 回答の配置パターンの内、
					// 入力条件として指定されているQの配置先を含まないものは表示しない
					return
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a.row[i] == j {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (a *Alds113a) recursive(i int) {
	if i == N {
		// 最後まで到達したときは配置成功
		// 結果を表示
		print()
		return
	}

	for j := 0; j < N; j++ {
		if a.row[j] == NOT_FREE || a.col[j] == NOT_FREE ||
			a.dpos[i+j] == NOT_FREE || a.dneg[i-j+N-1] == NOT_FREE {
			continue
		}

		// 1. 配置
		a.row[i], a.col[j], a.dpos[i+j], a.dneg[i-j+N-1] = j, NOT_FREE, NOT_FREE, NOT_FREE

		a.recursive(i + 1)

		// 1.の配置の状態をもとに戻して、次のマスに配置した時の配置パターンを検索
		a.row[i], a.col[j], a.dpos[i+j], a.dneg[i-j+N-1] = FREE, FREE, FREE, FREE
	}

	// 無理だった
}

func (a *Alds113a) main() {
	a.initialize()

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a.X[i][j] = false
		}
	}

	var k int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	k, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < k; i++ {
		var r, c int

		scanner.Scan()
		r, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		c, _ = strconv.Atoi(scanner.Text())

		a.X[r][c] = true
	}

	a.recursive(0)
}
