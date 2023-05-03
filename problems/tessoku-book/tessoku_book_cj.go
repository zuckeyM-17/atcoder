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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		fmt.Fprintln(out, bisearch(a, x))
	}
}

func bisearch(a []int, x int) int {
	l, r := 0, len(a)-1
	var p int
	for l <= r {
		p = (l + r) / 2
		if x <= a[p] {
			r = p - 1
		} else {
			l = p + 1
		}
	}
	return l
}
