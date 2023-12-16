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

	var n, q int
	fmt.Fscan(in, &n, &q)

	p := make([][]bool, n+1)
	for i := 1; i <= n; i++ {
		p[i] = make([]bool, n+1)
		var row string
		fmt.Fscan(in, &row)
		for j, c := range row {
			if c == 'B' {
				p[i][j] = true
			}
		}
	}

	lr, rl := make([][]int, n+2), make([][]int, n+2)

	for i := 0; i <= n+1; i++ {
		lr[i], rl[i] = make([]int, n+2), make([]int, n+2)
	}
	for i := 1; i <= n; i++ {
		lr[i], rl[i] = make([]int, n+2), make([]int, n+2)
		for j := 1; j <= n; j++ {
			if p[i][j] {
				lr[i][j] = lr[i][j-1] + 1
			}
			if p[i][n-j] {
				rl[i][n-j] = rl[i][n-j+1] + 1
			}
		}
	}

	lt, rt, lb, rb := make([][]int, n+2), make([][]int, n+2), make([][]int, n+2), make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		lt[i], rt[i], lb[i], rb[i] = make([]int, n+2), make([]int, n+2), make([]int, n+2), make([]int, n+2)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			lt[i][j] = lt[i-1][j] + lr[i][j]
			lb[n-1-i][j] = lb[n-i][j] + lr[n-1-i][j]
			rt[i][n-1-j] = rt[i][n-j] + rl[i][n-1-j]
			rb[n-1-i][n-1-j] = rb[n-i][n-j] + rl[n-1-i][n-1-j]
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, lt[i])
	}

	fmt.Fprintln(out)
}
