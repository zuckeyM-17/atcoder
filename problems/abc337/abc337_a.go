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
	var x, y int
	for i := 0; i < n; i++ {
		var xi, yi int
		fmt.Fscan(in, &xi, &yi)
		x += xi
		y += yi
	}

	if x == y {
		fmt.Fprintln(out, "Draw")
	} else if x > y {
		fmt.Fprintln(out, "Takahashi")
	} else {
		fmt.Fprintln(out, "Aoki")
	}
}
