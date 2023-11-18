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

	var n, q int
	fmt.Fscan(in, &n, &q)

	var s string
	fmt.Fscan(in, &s)

	sum := make([]int, n+1)

	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			sum[i] = sum[i-1] + 1
		} else {
			sum[i] = sum[i-1]
		}
	}
	sum[n] = sum[n-1]

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)

		fmt.Fprintln(out, sum[r-1]-sum[l-1])
	}
}
