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

	var a, b, c, d int

	s := make([][]bool, 101)
	for i := range s {
		s[i] = make([]bool, 101)
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		for j := a + 1; j <= b; j++ {
			for k := c + 1; k <= d; k++ {
				s[j][k] = true
			}
		}
	}

	ans := 0

	for r := range s {
		for i := range s[r] {
			if s[r][i] {
				ans += 1
			}
		}
	}

	fmt.Fprintln(out, ans)
}
