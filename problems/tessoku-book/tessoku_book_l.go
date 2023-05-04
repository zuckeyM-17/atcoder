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

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l, r := 0, 1_000_000_000

	var p int
	for l < r {
		p = (l + r) / 2
		if check(a, p, k) {
			r = p
		} else {
			l = p + 1
		}
	}

	fmt.Fprintln(out, p)
}

func check(a []int, p, k int) bool {
	var sum int
	for _, v := range a {
		sum += p / v
	}
	return sum >= k
}
