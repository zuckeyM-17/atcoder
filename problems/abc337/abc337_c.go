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

	a := map[int]int{}

	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		a[s] = i + 1
	}

	fmt.Fprint(out, a[-1], " ")
	next := a[-1]

	for i := 1; i < n; i++ {
		if i == n-1 {
			fmt.Fprintln(out, a[next])
		} else {
			fmt.Fprint(out, a[next], " ")
		}
		next = a[next]
	}
}
