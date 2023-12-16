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

	l, a := make([]int, 1000_001), make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		l[a[i]]++
	}

	sums := make([]int, 1000_001)
	sums[1000_000] = l[1000_000] * 1000_000
	for i := 999_999; i > 0; i-- {
		sums[i] = sums[i+1] + l[i]*i
	}

	for i := 1; i <= n; i++ {
		if a[i]+1 > 1000_000 {
			fmt.Fprintln(out, 0)
			continue
		}
		if i == n {
			fmt.Fprintln(out, sums[a[i]+1])
		} else {
			fmt.Fprint(out, sums[a[i]+1], " ")
		}
	}
}
