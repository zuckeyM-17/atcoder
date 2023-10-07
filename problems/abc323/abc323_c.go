package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type problem struct {
	score int
	order int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	problems := make([]problem, m)
	a := make([]int, m+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i+1])
		problems[i] = problem{a[i+1], i + 1}
	}

	sort.Slice(problems, func(i, j int) bool {
		return problems[i].score > problems[j].score
	})

	results := make([][]bool, n+1)
	scores := make([]int, n+1)
	maximum := -1
	for i := 1; i <= n; i++ {
		var s string
		results[i] = make([]bool, m+1)
		fmt.Fscan(in, &s)

		for j := 1; j <= m; j++ {
			if s[j-1] == 'x' {
				results[i][j] = false
			} else {
				results[i][j] = true
				scores[i] += a[j]
			}
		}
		scores[i] += i
		maximum = max(maximum, scores[i])
	}

	for i := 1; i <= n; i++ {
		if scores[i] == maximum {
			fmt.Fprintln(out, 0)
			continue
		}

		diff := maximum - scores[i]

		count := 0
		for _, p := range problems {
			if results[i][p.order] {
				continue
			}

			count++
			diff -= p.score

			if diff < 0 {
				break
			}
		}

		fmt.Fprintln(out, count)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
