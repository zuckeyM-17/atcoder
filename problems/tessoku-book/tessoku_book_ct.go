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

	var str string
	fmt.Fscan(in, &str)

	s := []rune(str)

	dp := [1009][1009]int{}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for length := 2; length <= n-1; length++ {
		for l := 0; l < n-length; l++ {
			r := l + length

			if s[l] == s[r] {
				dp[l][r] = max(max(dp[l][r-1], dp[l+1][r]), dp[l+1][r-1]+2)
			} else {
				dp[l][r] = max(dp[l][r-1], dp[l+1][r])
			}
		}
	}

	fmt.Fprintln(out, dp[0][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
