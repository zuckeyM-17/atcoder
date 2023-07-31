package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	var m, y, z int

	for i := 1; i <= n; i++ {
		m += abs(x[i])
		y += x[i] * x[i]
		z = max(z, abs(x[i]))
	}

	fmt.Fprintln(out, m)
	fmt.Fprintln(out, math.Sqrt(float64(y)))
	fmt.Fprintln(out, z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
