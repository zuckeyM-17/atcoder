package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func minCost(n, s, m, l int) int {
	maxEggs := 130
	cost := make([]int, maxEggs+1)

	for i := range cost {
		cost[i] = math.MaxInt32
	}

	cost[0] = 0

	for eggs := 1; eggs <= maxEggs; eggs++ {
		if eggs-6 >= 0 {
			cost[eggs] = min(cost[eggs], cost[eggs-6]+s)
		}
		if eggs-8 >= 0 {
			cost[eggs] = min(cost[eggs], cost[eggs-8]+m)
		}
		if eggs-12 >= 0 {
			cost[eggs] = min(cost[eggs], cost[eggs-12]+l)
		}
	}
	minCost := math.MaxInt32
	for i := n; i < maxEggs; i++ {
		minCost = min(minCost, cost[i])
	}

	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, s, m, l int
	fmt.Fscan(in, &n, &s, &m, &l)

	fmt.Fprintln(out, minCost(n, s, m, l))
}
