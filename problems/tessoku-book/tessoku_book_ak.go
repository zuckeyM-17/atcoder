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

	var n, m, b int
	fmt.Fscan(in, &n, &m, &b)

	ans := b * n * m
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans += a * m
	}

	for i := 1; i <= m; i++ {
		var c int
		fmt.Fscan(in, &c)
		ans += c * n
	}

	fmt.Fprintln(out, ans)
}
