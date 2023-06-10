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

	a := n / 5
	r := n % 5

	if r > 2 {
		a++
	}

	fmt.Fprintln(out, a*5)
}
