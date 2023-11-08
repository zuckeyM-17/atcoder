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

	var N, K int
	fmt.Fscan(in, &N, &K)
	A, B := make([]int, N), make([]int, N)

	minA, minB := 0, 0
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Fscan(in, &A[i], &B[i])
		minA = min(minA, a)
		minB = min(minB, b)
	}

	ans := 0

	for i := minA; i <= 100; i++ {
		for j := minB; j <= 100; j++ {
			cnt := 0
			for k := 0; k < N; k++ {
				if i <= A[k] && A[k] <= i+K && j <= B[k] && B[k] <= j+K {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	fmt.Fprintln(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
