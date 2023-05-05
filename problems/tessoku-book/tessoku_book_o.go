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

	b := make([]int, n)
	copy(b, a)

	sort.Ints(b)

	m := make(map[int]bool)
	uniq := []int{}
	for _, v := range b {
		if !m[v] {
			m[v] = true
			uniq = append(uniq, v)
		}
	}

	x := []int{}
	for _, v := range a {
		l, r := 0, len(uniq)
		for l < r {
			mid := (l + r) / 2
			if uniq[mid] < v {
				l = mid + 1
			} else {
				r = mid
			}
		}
		x = append(x, l+1)
	}

	for i, v := range x {
		fmt.Fprint(out, v)
		if i < len(x)-1 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out, "")
		}
	}
}
