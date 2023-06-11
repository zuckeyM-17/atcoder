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

	var a, b int
	fmt.Fscan(in, &a, &b)

	p, ans := a, 1
	for i := 0; i < 60; i++ {
		w := 1 << i
		if (b/w)%2 == 1 {
			ans = (ans * p) % 1000000007
		}
		p = (p * p) % 1000000007
	}

	fmt.Fprintln(out, ans)
}
