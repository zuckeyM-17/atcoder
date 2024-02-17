package main

import (
	"bufio"
	"fmt"
	"os"
)

type step struct {
	h, w int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, n int
	var t string
	fmt.Fscan(in, &h, &w, &n)
	fmt.Fscan(in, &t)

	steps := make([]step, n+1)
	for i := 1; i <= n; i++ {
		switch t[i-1] {
		case 'U':
			steps[i] = step{-1, 0}
		case 'D':
			steps[i] = step{1, 0}
		case 'L':
			steps[i] = step{0, -1}
		case 'R':
			steps[i] = step{0, 1}
		}
	}

	a := make([][]rune, h+1)
	for i := 1; i <= h; i++ {
		a[i] = make([]rune, w+1)
		var s string
		fmt.Fscan(in, &s)
		for j, c := range s {
			a[i][j+1] = c
		}
	}

	cnt := 0
	for i := 2; i <= h-1; i++ {
		for j := 2; j <= w-1; j++ {
			if a[i][j] == '#' {
				continue
			}

			cur := step{i, j}
			ok := true
			for k := 1; k <= n; k++ {
				if a[cur.h+steps[k].h][cur.w+steps[k].w] == '#' {
					ok = false
					break
				} else {
					cur.h += steps[k].h
					cur.w += steps[k].w
				}
			}
			if ok {
				cnt++
			}
		}
	}

	fmt.Fprintln(out, cnt)
}
