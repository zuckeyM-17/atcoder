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

	var n, x int
	fmt.Fscan(in, &n, &x)

	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i])
	}

	var b []int
	for i := 0; i <= 100; i++ {
		b = append(a, i)
		sort.Ints(b)
		if sum(b) >= x {
			fmt.Fprintln(out, i)
			return
		}
	}

	fmt.Fprintln(out, -1)
}

func sum(a []int) int {
	ans := 0
	for i := 1; i < len(a)-1; i++ {
		ans += a[i]
	}
	return ans
}
