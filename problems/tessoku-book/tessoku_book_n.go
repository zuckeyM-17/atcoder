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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a, b, c, d := make([]int, n), make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}

	p, q := []int{}, []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			p = append(p, a[i]+b[j])
			q = append(q, c[i]+d[j])
		}
	}

	sort.Ints(q)

	for _, v := range p {
		if k-v < 0 {
			continue
		}
		l, r := 0, len(q)-1
		for l < r {
			m := (l + r) / 2
			if q[m] == k-v {
				fmt.Fprintln(out, "Yes")
				return
			}
			if q[m] < k-v {
				l = m + 1
			} else {
				r = m
			}
		}
	}

	fmt.Fprintln(out, "No")
}
