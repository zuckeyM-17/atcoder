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

	r := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			r[i] = 0
		} else {
			r[i] = r[i-1]
		}

		for r[i] < n-1 && a[r[i]+1]-a[i] <= k {
			r[i]++
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += r[i] - i
	}
	fmt.Fprintln(out, res)
}
