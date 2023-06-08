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
	oa, ob := a, b
	for a >= 1 && b >= 1 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	var cd int
	if a != 0 {
		cd = a
	} else {
		cd = b
	}

	fmt.Fprintln(out, oa*ob/cd)
}
