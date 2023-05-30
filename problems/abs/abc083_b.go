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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)

	ans := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}

		if a <= sum && sum <= b {
			ans += i
		}
	}

	fmt.Fprintln(out, ans)
}
