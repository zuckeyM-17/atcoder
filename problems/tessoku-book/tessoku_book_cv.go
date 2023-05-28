package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x, y := make([]float64, n), make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	dp := make([][]float64, (1<<n)+9)
	for i := 0; i < (1<<n)+9; i++ {
		dp[i] = make([]float64, n+9)
		for j := 0; j < n+9; j++ {
			dp[i][j] = 100_000_000_000
		}
	}

	dp[0][0] = 0

	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1e9 {
				continue
			}

			// j => k への移動を考える
			// k が訪問済みならスキップ
			for k := 0; k < n; k++ {
				if (i/1<<k)%2 == 1 {
					continue
				}

				// j => k への移動のコストを計算
				dist := math.Pow(math.Pow(x[j]-x[k], 2)+math.Pow(y[j]-y[k], 2), 0.5)
				dp[i|1<<k][k] = min(dp[i|1<<k][k], dp[i][j]+dist)
			}
		}
	}

	fmt.Fprintln(out, dp[(1<<n)-1][0])
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
