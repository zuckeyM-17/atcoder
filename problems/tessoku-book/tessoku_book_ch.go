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

	grid := make([][]int, 1502)
	for i := 0; i < 1502; i++ {
		grid[i] = make([]int, 1502)
	}

	var n int
	fmt.Fscan(in, &n)

	var a, b, c, d int

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c, &d)
		// fmt.Fprintln(out, a, b, c, d, grid[a], grid[c])
		grid[a][b]++
		grid[c][d]++
		grid[a][d]--
		grid[c][b]--
	}

	for i := 0; i < 1500; i++ {
		for j := 0; j < 1500; j++ {
			grid[i][j+1] += grid[i][j]
		}
	}

	for i := 0; i < 1500; i++ {
		for j := 0; j < 1500; j++ {
			grid[i+1][j] += grid[i][j]
		}
	}

	count := 0
	for i := 0; i < 1500; i++ {
		for j := 0; j < 1500; j++ {
			if grid[i][j] > 0 {
				count++
			}
		}
	}

	fmt.Fprintln(out, count)
}
