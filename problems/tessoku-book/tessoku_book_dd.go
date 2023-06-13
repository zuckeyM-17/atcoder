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

	ans := n/3 + n/5 + n/7 - n/15 - n/21 - n/35 + n/105

	fmt.Fprintln(out, ans)
}
