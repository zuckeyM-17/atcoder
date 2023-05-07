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

	h := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = abs(h[2] - h[1])
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+abs(h[i]-h[i-1]), dp[i-2]+abs(h[i]-h[i-2]))
	}

	places := []int{}
	i := n
	for {
		places = append(places, i)
		if i == 1 {
			break
		}

		if dp[i-1]+abs(h[i]-h[i-1]) == dp[i] {
			i -= 1
		} else {
			i -= 2
		}
	}

	sort.Slice(places, func(i, j int) bool {
		return places[i] < places[j]
	})

	fmt.Fprintln(out, len(places))
	for i := 0; i < len(places); i++ {
		fmt.Fprint(out, places[i])
		if i == len(places)-1 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 引数にとった整数の絶対値を返す
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
