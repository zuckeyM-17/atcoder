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

	for i := 0; i <= n-3; i++ {
		if s[i] == 'R' && s[i+1] == 'R' && s[i+2] == 'R' {
			fmt.Fprintln(out, "Yes")
			return
		}
		if s[i] == 'B' && s[i+1] == 'B' && s[i+2] == 'B' {
			fmt.Fprintln(out, "Yes")
			return
		}
	}

	fmt.Fprintln(out, "No")
}
