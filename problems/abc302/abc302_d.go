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

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)

	a, b := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)

	for {
		x, y := len(a), len(b)
		if x == 0 || y == 0 {
			fmt.Fprintln(out, -1)
			break
		}
		if abs(a[x-1]-b[y-1]) <= d {
			fmt.Fprintln(out, a[x-1]+b[y-1])
			break
		}

		if a[x-1] > b[y-1] {
			a = a[:x-1]
		} else {
			b = b[:y-1]
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
