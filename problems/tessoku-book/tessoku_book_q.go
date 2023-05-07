package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a, b := make([]int, n+1), make([]int, n+1)

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = a[2]
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+a[i], dp[i-2]+b[i])
	}

	rooms := []int{}
	i := n
	for {
		rooms = append(rooms, i)
		if i == 1 {
			break
		}

		if dp[i-1]+a[i] == dp[i] {
			i -= 1
		} else {
			i -= 2
		}
	}

	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i] < rooms[j]
	})

	fmt.Fprintln(out, len(rooms))
	for i := 0; i < len(rooms); i++ {
		fmt.Fprint(out, rooms[i])
		if i == len(rooms)-1 {
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, " ")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
