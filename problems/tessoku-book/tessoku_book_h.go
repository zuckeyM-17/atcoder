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

	var h, w, x, q, a, b, c, d int
	fmt.Fscan(in, &h, &w)

	sum := [][]int{}
	for i := 0; i <= h; i++ {
		sum = append(sum, make([]int, w+1))
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Fscan(in, &x)
			sum[i][j] = sum[i][j-1] + x
		}
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
			sum[j][i] += sum[j-1][i]
		}
	}

	fmt.Fscan(in, &q)

	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		fmt.Fprintln(out, sum[c][d]+sum[a-1][b-1]-sum[c][b-1]-sum[a-1][d])
	}
}
