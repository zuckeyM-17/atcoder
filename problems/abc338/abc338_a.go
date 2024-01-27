package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUpperCase(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	for i, r := range s {
		if i == 0 {
			if !isUpperCase(r) {
				fmt.Fprintln(out, "No")
				return
			}
		} else {
			if isUpperCase(r) {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
}
