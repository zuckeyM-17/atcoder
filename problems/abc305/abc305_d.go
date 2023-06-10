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
	a, fa := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if i == 0 {
			fa[i] = a[i]
		} else if i%2 == 0 {
			fa[i] = fa[i-1] + a[i] - a[i-1]
		} else {
			fa[i] = fa[i-1]
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, f(r, a, fa)-f(l, a, fa))
	}
}

func f(x int, a, fa []int) int {
	j := sort.Search(len(a), func(i int) bool { return a[i] > x })
	if j == len(a) {
		return fa[len(a)-1]
	}
	j--
	return fa[j] + (x-a[j])*(fa[j+1]-fa[j])/(a[j+1]-a[j])
}
