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

	var n, x, y, q, a, b, c, d int

	fmt.Fscan(in, &n)

	var dp [1501][1501]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		dp[x][y] = dp[x][y] + 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp); j++ {
			dp[i][j] = dp[i][j] + dp[i-1][j]
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp); j++ {
			dp[i][j] = dp[i][j] + dp[i][j-1]
		}
	}

	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		fmt.Fprintln(out, dp[c][d]-dp[a-1][d]-dp[c][b-1]+dp[a-1][b-1])
	}
}
