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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	p := enumerate(a[:n/2])
	q := enumerate(a[n/2:])

	sort.Ints(q)

	for _, v := range p {
		if k-v < 0 {
			continue
		}
		l, r := 0, len(q)-1
		for l <= r {
			m := (l + r) / 2
			if q[m] == k-v {
				fmt.Fprintln(out, "Yes")
				return
			}
			if q[m] < k-v {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	fmt.Fprintln(out, "No")
}

func enumerate(a []int) []int {
	var sumList []int
	for i := 0; i < 1<<len(a); i++ {
		sum := 0
		for j := 0; j < len(a); j++ {
			w := 1 << j
			if (i/w)%2 == 1 {
				sum += a[j]
			}
		}
		sumList = append(sumList, sum)
	}
	return sumList
}
