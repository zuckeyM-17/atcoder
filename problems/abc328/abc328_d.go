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

	var s string
	fmt.Fscan(in, &s)

	cur := 0
	for cur+2 <= len(s)-1 && len(s) > 2 {
		if s[cur] == 'A' && s[cur+1] == 'B' && s[cur+2] == 'C' {
			s = s[:cur] + s[cur+3:]
			if cur > 1 {
				cur = cur - 2
			} else if cur > 0 {
				cur = cur - 1
			}
		} else {
			cur++
		}
	}

	fmt.Fprintln(out, s)
}
