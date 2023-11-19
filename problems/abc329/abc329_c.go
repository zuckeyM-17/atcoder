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

	var s string
	fmt.Fscan(in, &s)

	m := map[string]bool{}

	l, r := 0, 1

	m[string(s[0])] = true
	for l < n && r < n {
		if s[l] == s[r] {
			m[string(s[l:r+1])] = true
			r++
			continue
		} else {
			l = r
		}
	}

	fmt.Fprintln(out, len(m))
}
