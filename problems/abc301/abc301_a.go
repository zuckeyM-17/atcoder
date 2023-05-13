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

	var s string
	fmt.Fscan(in, &s)

	t := 0
	for _, c := range s {
		if c == 'T' {
			t++
		}
	}

	a := n - t
	if a == t {
		if s[len(s)-1] == 'T' {
			fmt.Fprintln(out, "A")
		} else {
			fmt.Fprintln(out, "T")
		}
	} else if a > t {
		fmt.Fprintln(out, "A")
	} else {
		fmt.Fprintln(out, "T")
	}
}
