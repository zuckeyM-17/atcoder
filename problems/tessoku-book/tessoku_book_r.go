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

	var n, s int
	fmt.Fscan(in, &n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			} else if j >= a[i-1] && dp[i-1][j-a[i-1]] {
				dp[i][j] = true
			}
		}
	}

	if dp[n][s] {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
