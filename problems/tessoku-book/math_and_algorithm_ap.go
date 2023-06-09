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

	a, b := 1, 1

	for i := 3; i <= n; i++ {
		a, b = b, (a+b)%1000_000_007
	}

	fmt.Fprintln(out, b)
}
