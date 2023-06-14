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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
	}

	dp := make([]bool, n+1)

	for i := 0; i <= n; i++ {
		for _, v := range a {
			if i >= v && !dp[i-v] {
				dp[i] = true
			}
		}
	}

	if dp[n] {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
