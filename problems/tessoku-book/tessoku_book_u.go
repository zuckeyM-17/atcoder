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

	var n int
	fmt.Fscan(in, &n)

	fmt.Fprintln(out, n)

	p, a := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		fmt.Fscan(in, &a[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for LEN := n - 2; LEN >= 0; LEN-- {
		for l := 1; l <= n-LEN; l++ {
			fmt.Fprintln(out, dp, l, LEN)
			r := l + LEN

			var score1, score2 int
			if l <= p[l-1] && p[l-1] <= r {
				score1 = a[l-1]
			}

			if l <= p[r+1] && p[r+1] <= r {
				score2 = a[r+1]
			}

			if l == 1 {
				dp[l][r] = dp[l][r+1] + score2
			} else if r == n {
				dp[l][r] = dp[l-1][r] + score1
			} else {
				dp[l][r] = max(dp[l-1][r]+score1, dp[l][r+1]+score2)
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans = max(ans, dp[i][i])
	}

	fmt.Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
