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

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := make([]int, n+1)

	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
	}

	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			r[i] = 0
		} else {
			r[i] = r[i-1]
		}
		for r[i] < n && s[r[i]+1]-s[i-1] <= k {
			r[i]++
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res += r[i] - i + 1
	}
	fmt.Fprintln(out, res)
}
