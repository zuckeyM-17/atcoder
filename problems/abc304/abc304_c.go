package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, d int
	fmt.Fscan(in, &n, &d)

	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	infected := make([]bool, n)
	infected[0] = true
	neighbors := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if distance(x[i], y[i], x[j], y[j]) <= float64(d) {
				neighbors[i] = append(neighbors[i], j)
				if infected[j] {
					infected[i] = true
				}
			}
		}
	}
	count := 0
	for count < 2000 {
		count++
		for i := 1; i < n; i++ {
			if infected[i] {
				continue
			} else {
				for _, v := range neighbors[i] {
					if infected[v] {
						infected[i] = true
						break
					}
				}
			}
		}
	}

	for _, v := range infected {
		if v {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}
