package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wt := bufio.NewWriter(os.Stdout)
	defer wt.Flush()

	var t, n int
	fmt.Scanf("%d", &t)
	fmt.Scanf("%d", &n)

	b := make([]int, t+1)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)
		b[l]++
		b[r]--
	}

	sum := []int{0}
	for i := 0; i < len(b)-1; i++ {
		sum = append(sum, sum[i]+b[i])
	}

	for i := 1; i < len(sum); i++ {
		fmt.Fprintln(wt, sum[i])
	}
}
