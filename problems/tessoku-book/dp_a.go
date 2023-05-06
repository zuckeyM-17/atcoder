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

	h := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
	}

	dp := make([]int, n+1)
	dp[1] = 0

	fmt.Fprintln(out, dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
