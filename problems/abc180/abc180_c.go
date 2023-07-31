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

	a := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			j := n / i
			if i == j {
				a = append(a, i)
			} else {
				a = append(a, i)
				a = append(a, j)
			}
		}
	}

	sort.Ints(a)

	for _, v := range a {
		fmt.Fprintln(out, v)
	}
}
