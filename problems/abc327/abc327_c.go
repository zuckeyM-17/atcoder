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

	a := make([][]int, 9)
	for i := range a {
		a[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for i := 0; i < 9; i++ {
		if !isAllUnique(a[i]) {
			fmt.Fprintln(out, "No")
			return
		}
	}

	for i := 0; i < 9; i++ {
		b := make([]int, 9)
		for j := 0; j < 9; j++ {
			b[j] = a[j][i]
		}
		if !isAllUnique(b) {
			fmt.Fprintln(out, "No")
			return
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			b := make([]int, 9)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					b[k*3+l] = a[i+k][j+l]
				}
			}
			if !isAllUnique(b) {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
}

func isAllUnique(a []int) bool {
	m := make(map[int]bool)
	for _, v := range a {
		if m[v] {
			return false
		}
		m[v] = true
	}
	return true
}
