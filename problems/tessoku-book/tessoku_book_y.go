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

	var h, w int
	fmt.Fscan(in, &h, &w)

	c, dp := make([][]string, h+1), make([][]int, h+1)
	for i := 0; i <= h; i++ {
		c[i] = make([]string, w+1)
		dp[i] = make([]int, w+1)
		if i > 0 {
			var s string
			fmt.Fscan(in, &s)
			for j, v := range s {
				c[i][j+1] = string(v)
			}
		}
	}

	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}

			if i == 1 && j == 1 {
				dp[i][j] = 1
				continue
			}

			if c[i][j] == "#" {
				dp[i][j] = 0
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Fprintln(out, dp[h][w])
}
