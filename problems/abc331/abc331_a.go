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

	var M, D int
	fmt.Fscan(in, &M, &D)
	var y, m, d int
	fmt.Fscan(in, &y, &m, &d)

	if d == D {
		if m == M {
			fmt.Fprintln(out, y+1, 1, 1)
			return
		} else {
			fmt.Fprintln(out, y, m+1, 1)
			return
		}
	}

	fmt.Fprintln(out, y, m, d+1)
}
