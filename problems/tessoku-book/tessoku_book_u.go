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
	fmt.Fscan(in, n)

	var p, a []int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		fmt.Fscan(in, &a[i])
	}

	fmt.Fprintln(out)
}
