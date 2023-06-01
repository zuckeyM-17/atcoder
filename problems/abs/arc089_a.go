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

	t, x, y := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t[i], &x[i], &y[i])
		x, y := abs(x[i]-x[i-1]), abs(y[i]-y[i-1])
		if x+y > t[i]-t[i-1] || (x+y+t[i]-t[i-1])%2 != 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
