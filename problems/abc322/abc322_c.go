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

	var n, m int
	fmt.Fscan(in, &n, &m)

	a := make([]int, m+1)

	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}

	cur := 1
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, a[cur]-i)
		if a[cur] == i {
			cur++
		}
	}
}
