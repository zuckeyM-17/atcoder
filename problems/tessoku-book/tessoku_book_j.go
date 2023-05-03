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

	a, p, q := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[i] = x
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			p[i] = a[i]
			q[n-i-1] = a[n-i-1]
			continue
		}
		p[i] = max(p[i-1], a[i])
		q[n-i-1] = max(q[n-i], a[n-i-1])
	}

	var d int
	fmt.Fscan(in, &d)
	for i := 0; i < d; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		res := max(p[l-2], q[r])
		fmt.Fprintln(out, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
