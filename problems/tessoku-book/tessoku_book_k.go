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

	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l, r := 0, n-1
	p := culcP(l, r)
	for {
		if a[p] == x {
			fmt.Fprintln(out, p+1)
			return
		} else if a[p] < x {
			l = p + 1
			p = culcP(l, r)
		} else {
			r = p - 1
			p = culcP(l, r)
		}
	}
}

func culcP(l, r int) int {
	return (l + r) / 2
}
