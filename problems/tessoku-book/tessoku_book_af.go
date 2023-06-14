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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)

	dp := make([]bool, n+1)

	for i := 0; i <= n; i++ {
		if i >= a && dp[i-a] == false {
			dp[i] = true
		} else if i >= b && dp[i-b] == false {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}

	if dp[n] {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
