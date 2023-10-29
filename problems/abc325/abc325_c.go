package main

import (
	"bufio"
	"fmt"
	"os"
)

type square struct {
	s       rune
	checked bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([][]square, h)
	for i := 0; i < h; i++ {
		var str string
		fmt.Fscan(in, &str)
		s[i] = make([]square, w)
		for j, r := range str {
			s[i][j] = square{s: r, checked: false}
		}
	}

	num := 1
	for i, row := range s {
		for j, sq := range row {
			if sq.s == '.' || sq.checked {
				continue
			}

		}
	}

	fmt.Fprintln(out, num)
}

func check(i, j, h, w int, s *[][]square) {
	adjs := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, adj := range adjs {
		adjI, adjJ := i+adj[0], j+adj[1]
		if adjI < 0 || adjJ < 0 || adjI >= h || adjJ >= w {
			continue
		}
		if s[adjI][adjJ].s == '.' {
			s[adjI][adjJ].checked = true
			continue
		}

		if s[adjI][adjJ].s == '#' {

		}
	}
}
