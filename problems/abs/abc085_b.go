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

	d := make([]int, n)
	m := make(map[int]bool)
	uniq := []int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
		if !m[d[i]] {
			m[d[i]] = true
			uniq = append(uniq, d[i])
		}
	}

	fmt.Fprintln(out, len(uniq))
}
