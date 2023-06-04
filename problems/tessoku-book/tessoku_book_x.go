package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	sort.Ints(dp)
	fmt.Fprintln(out, dp[len(dp)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
