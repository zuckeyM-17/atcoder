package main

import (
	"bufio"
	"fmt"
	"os"
)

type cost struct {
	c int
	a []int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, p int
	fmt.Fscan(in, &n, &k, &p)

	c := make([]int, n+1)
	a := make([][]int, n+1)
	sum := make([]int, k+1)

	for i := 1; i <= n; i++ {
		a[i] = make([]int, k+1)
	}

	ans := 1000_000_000
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
		for j := 1; j <= k; j++ {
			fmt.Fscan(in, &a[i][j])
			if allAbove(a[i], p) {
				ans = min(c[i], ans)
			}
			sum[j] += a[i][j]
		}
	}

	if ans != 1000_000_000 {
		fmt.Fprintln(out, ans)
		return
	}

	for _, v := range sum {
		if v > p {
			continue
		} else {
			fmt.Fprintln(out, -1)
			return
		}
	}

	dp := make([][]cost, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]cost, 32)
		for j := 0; j < 32; j++ {
			dp[i][j] = cost{1000_000_000, make([]int, k+1)}
		}
	}

	dp[0][0] = cost{0, make([]int, k+1)}
	for i := 1; i <= n; i++ {
		for j := 0; j < 32; j++ {
			al := make([]int, k+1)

			for l := 1; l <= n; l++ {

			}

			dp[i][j] = minCost(dp[i][j], dp[i-1][j])
			dp[i][v] = minCost(dp[i][v], dp[i-1][j])
		}
	}
	fmt.Fprintln(out, sum)
}

func minCost(a, b cost) cost {
	if a.c < b.c {
		return a
	}
	return b
}

func plusSlice(a []int, b []int) []int {
	c := make([]int, len(a))
	for i := 1; i <= len(a); i++ {
		c[i] = a[i] + b[i]
	}

	return c
}

func allAbove(a []int, p int) bool {
	for i := 1; i <= len(a); i++ {
		if a[i] <= p {
			return false
		}
	}
	return true
}
