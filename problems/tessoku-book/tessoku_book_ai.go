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

	dp := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &dp[n][i])
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			if i%2 == 1 {
				dp[i][j] = max([]int{dp[i+1][j], dp[i+1][j+1]})
			}
			if i%2 == 0 {
				dp[i][j] = min([]int{dp[i+1][j], dp[i+1][j+1]})
			}
		}
	}

	fmt.Fprintln(out, dp[1][1])
}

// slice から最大値を返す
func max(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func min(slice []int) int {
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}
