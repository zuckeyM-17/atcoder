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

	r := make([][]int, 18)
	p := make([]int, 18)

	for i := 0; i < 18; i++ {
		r[i] = make([]int, 10)
		if i == 0 {
			p[0] = 1
		} else {
			p[i] = 10 * p[i-1]
		}
	}

	for i := 0; i <= 14; i++ {
		c := (n / p[i]) % 10
		for j := 0; j < 10; j++ {
			if j < c {
				r[i][j] = (n/p[i+1])*p[i] + p[i]
			} else if j == c {
				r[i][j] = (n/p[i+1])*p[i] + n%p[i] + 1
			} else {
				r[i][j] = (n / p[i+1]) * p[i]
			}
		}
	}

	ans := 0
	for i := 0; i < 15; i++ {
		for j := 0; j < 10; j++ {
			ans += j * r[i][j]
		}
	}

	fmt.Fprintln(out, ans)
}
