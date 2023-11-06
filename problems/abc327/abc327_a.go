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
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)

	var t rune
	for i, c := range s {
		if i == 0 {
			continue
		}
		t = rune(s[i-1])
		if t == 'a' && c == 'b' {
			fmt.Fprintln(out, "Yes")
			return
		} else if t == 'b' && c == 'a' {
			fmt.Fprintln(out, "Yes")
			return
		}
	}
	fmt.Fprintln(out, "No")
}
