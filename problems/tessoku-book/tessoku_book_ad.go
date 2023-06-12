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

	m := 1000000007
	var n, r int
	fmt.Fscan(in, &n, &r)

	a, b := 1, 1

	for i := r + 1; i <= n; i++ {
		a *= i
		a %= m
	}

	for i := 1; i <= n-r; i++ {
		b *= i
		b %= m
	}

	fmt.Fprintln(out, a*power(b, m-2, m)%m)
}

func power(a, b, m int) int {
	p, ans := a, 1
	for i := 0; i < 30; i++ {
		w := 1 << i
		if (b/w)%2 == 1 {
			ans = (ans * p) % m
		}
		p = (p * p) % m
	}
	return ans
}
