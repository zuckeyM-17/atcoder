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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := []int{}
	for i := 0; i < n; i++ {
		pos := sort.SearchInts(l, a[i])
		if pos == len(l) {
			l = append(l, a[i])
		} else {
			l[pos] = a[i]
		}
	}

	fmt.Fprintln(out, len(l))
}
