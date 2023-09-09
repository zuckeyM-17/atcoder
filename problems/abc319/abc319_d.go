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

	var n, m int
	fmt.Fscan(in, &n, &m)

	l := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
		sum += l[i]
	}

	w := sum / m

	for {
		rows := m
		cur := 0
		for i := 0; i < n; i++ {
			cur += l[i]

			if i == n-1 && cur <= w {
				break
			}

			if cur > w {
				rows -= 1
				cur = l[i]
			} else if cur == w {
				rows -= 1
				cur = 0
			} else {
				cur += 1
			}
			if rows == 0 {
				break
			}
		}
		if rows > 0 {
			break
		}
		w += 1
	}

	fmt.Fprintln(out, w)
}
