package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N        = 8
	FREE     = -1
	NOT_FREE = -2
)

var row, col [N]int
var dpos, dneg [2*N - 1]int
var X [N][N]bool

func initialize() {
	for i := 0; i < N; i++ {
		row[i], col[i] = FREE, FREE
	}
	for i := 0; i < 2*N-1; i++ {
		dpos[i], dneg[i] = FREE, FREE
	}
}

func print() {
	// 回答となる配置パターンの数だけprintがコールされる
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if X[i][j] {
				if row[i] != j {
					// 回答の配置パターンの内、
					// 入力条件として指定されているQの配置先を含まないものは表示しない
					return
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if row[i] == j {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func recursive(i int) {
	if i == N {
		// 最後まで到達したときは配置成功
		// 結果を表示
		print()
		return
	}

	for j := 0; j < N; j++ {
		if row[j] == NOT_FREE || col[j] == NOT_FREE ||
			dpos[i+j] == NOT_FREE || dneg[i-j+N-1] == NOT_FREE {
			continue
		}

		// 1. 配置
		row[i], col[j], dpos[i+j], dneg[i-j+N-1] = j, NOT_FREE, NOT_FREE, NOT_FREE

		recursive(i + 1)

		// 1.の配置の状態をもとに戻して、次のマスに配置した時の配置パターンを検索
		row[i], col[j], dpos[i+j], dneg[i-j+N-1] = FREE, FREE, FREE, FREE
	}

	// 無理だった
}

func main() {
	initialize()

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			X[i][j] = false
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

		X[r][c] = true
	}

	recursive(0)
}

