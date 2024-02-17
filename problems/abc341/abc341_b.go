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

	a, s, t := make([]int, n+1), make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 1; i <= n-1; i++ {
		fmt.Fscan(in, &s[i], &t[i])
	}

	for i := 1; i <= n-1; i++ {
		a[i+1] = a[i+1] + (a[i] / s[i] * t[i])
	}

	fmt.Fprintln(out, a[n])
}
