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

	var n, m, p int
	fmt.Fscan(in, &n, &m, &p)

	t := n - m

	if t < 0 {
		fmt.Fprintln(out, 0)
		return
	}

	a := t / p

	fmt.Fprintln(out, a+1)
}
