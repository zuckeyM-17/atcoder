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

	p, a := make([]int, n+2), make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		fmt.Fscan(in, &a[i])
	}

	dp := [2009][2009]int{}

	dp[1][n] = 0

	for length := n - 2; length >= 0; length-- {
		for l := 1; l <= n-length; l++ {
			r := l + length

			score1, score2 := 0, 0
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
				dp[l][r] = max(dp[l][r+1]+score2, dp[l-1][r]+score1)
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
