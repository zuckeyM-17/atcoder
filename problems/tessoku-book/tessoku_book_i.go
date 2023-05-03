package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, n, a, b, c, d int
	fmt.Fscan(in, &h, &w, &n)

	// h * w のマス目を作成
	m := make([][]int, h+1)
	for i := 0; i < h+1; i++ {
		m[i] = make([]int, w+1)
	}

	// n 回の操作を行う
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		m[a-1][b-1] = m[a-1][b-1] + 1
		m[a-1][d] = m[a-1][d] - 1
		m[c][b-1] = m[c][b-1] - 1
		m[c][d] = m[c][d] + 1
	}

	// 二次元累積和を求める
	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			m[i][j] = m[i][j] + m[i][j-1]
		}
	}
	for i := 1; i < h; i++ {
		for j := 0; j < w; j++ {
			m[i][j] = m[i][j] + m[i-1][j]
		}
	}

	// 結果を出力
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(out, m[i][j])
			if j < w-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
