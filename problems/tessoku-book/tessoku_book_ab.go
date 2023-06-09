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

	ans := 0
	for i := 0; i < n; i++ {
		var (
			t string
			a int
		)
		fmt.Fscan(in, &t, &a)

		if t == "+" {
			ans += a
		} else if t == "-" {
			ans -= a
		} else if t == "*" {
			ans *= a
		}
		if ans < 0 {
			ans += 100_00
		}
		ans %= 100_00
		fmt.Fprintln(out, ans)
	}
}
