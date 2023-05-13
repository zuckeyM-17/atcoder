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

	var (
		s string
		n int
	)

	fmt.Fscan(in, &s)
	fmt.Fscan(in, &n)

	ans := 0
	for i, c := range s {
		idx := len(s) - i - 1
		if c == '1' {
			ans += 1 << idx
		}
	}
	if ans > n {
		fmt.Fprintln(out, "-1")
		return
	}

	for i, c := range s {
		if c != '?' {
			continue
		}
		idx := len(s) - i - 1
		add := 1 << idx
		if ans+add <= n {
			ans += add
		}
	}

	fmt.Fprintln(out, ans)
}
