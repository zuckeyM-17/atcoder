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

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice, bob := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	fmt.Fprintln(out, alice-bob)
}
