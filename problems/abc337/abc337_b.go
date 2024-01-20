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

	if s[0] != 'A' && s[0] != 'B' && s[0] != 'C' {
		fmt.Fprintln(out, "No")
		return
	}

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			continue
		} else if s[i-1] == 'A' && s[i] == 'B' {
			continue
		} else if s[i-1] == 'B' && s[i] == 'C' {
			continue
		} else if s[i-1] == 'A' && s[i] == 'C' {
			continue
		} else {
			fmt.Fprintln(out, "No")
			return
		}
	}

	fmt.Fprintln(out, "Yes")
}
