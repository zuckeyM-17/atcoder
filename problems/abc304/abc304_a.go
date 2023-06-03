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

	s, a := make([]string, n), make([]int, n)

	min, first := 1000000009, -1
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i], &a[i])
		if a[i] < min {
			min = a[i]
			first = i
		}
	}

	t := []string{}
	t = append(t, s[first:]...)
	t = append(t, s[:first]...)

	for _, v := range t {
		fmt.Fprintln(out, v)
	}
}
