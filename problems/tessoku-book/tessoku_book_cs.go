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

	var s, t string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &t)

	S, T := len(s), len(t)

	dp := make([][]int, S+1)
	for i := 0; i <= S; i++ {
		dp[i] = make([]int, T+1)
	}

	for i := 0; i <= S; i++ {
		for j := 0; j <= T; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			cost := 0
			if s[i-1] != t[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost,
			)
		}
	}

	fmt.Fprintln(out, dp[S][T])
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
