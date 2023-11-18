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

	var n, x int
	fmt.Fscan(in, &n, &x)

	ans := 0
	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		if s <= x {
			ans += s
		}
	}

	fmt.Fprintln(out, ans)
}
