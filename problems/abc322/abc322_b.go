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

	var s, t string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &t)

	if t[:len(s)] == s && t[len(t)-len(s):] == s {
		fmt.Fprintln(out, 0)
	} else if t[:len(s)] == s {
		fmt.Fprintln(out, 1)
	} else if t[len(t)-len(s):] == s {
		fmt.Fprintln(out, 2)
	} else {
		fmt.Fprintln(out, 3)
	}
}
