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

	var n int
	fmt.Fscan(in, &n)

	w := make([]int, n)
	x := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i], &x[i])
	}

	ans := 0
	for i := 0; i < 24; i++ {
		num := 0
		for j, hour := range x {
			if (hour+i)%24 >= 9 && (hour+i)%24 <= 17 {
				num += w[j]
			}
		}
		ans = max(ans, num)
	}

	fmt.Fprintln(out, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
