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

	var a, b int
	fmt.Fscan(in, &a, &b)

	for a >= 1 && b >= 1 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a != 0 {
		fmt.Fprintln(out, a)
		return
	}
	fmt.Fprintln(out, b)
}
