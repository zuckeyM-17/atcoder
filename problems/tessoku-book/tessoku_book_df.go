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

	var n, h, w int
	fmt.Fscan(in, &n, &h, &w)

	a, b := make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	xorSum := 0

	for i := 1; i <= n; i++ {
		xorSum = xorSum ^ (a[i] - 1)
	}

	for i := 1; i <= n; i++ {
		xorSum = xorSum ^ (b[i] - 1)
	}

	if xorSum != 0 {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
