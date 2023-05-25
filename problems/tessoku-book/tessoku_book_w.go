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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 1; i < m; i++ {
		for j := 0; j < 1<<n; j++ {
			already := make([]int, n)
			for k := 0; k < n; k++ {
				if j&1<<k != 0 {
					already[k] = 1
				}

				v := 0
				for k := 0; k < n; k++ {
					if already[k] == 1 || a[i][k] == 1 {
						v += 1 << k
					}
				}

				dp[i][j] = min(dp[i][j], dp[i-1][j])
				dp[i][v] = min(dp[i][v], dp[i-1][j]+1)
			}
		}
	}

	if dp[m-1][(1<<n)-1] == 1<<60 {
		fmt.Fprintln(out, -1)
		return
	}

	fmt.Fprintln(out, dp[m-1][(1<<n)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
