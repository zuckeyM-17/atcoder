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

	var ans int
	if a%b > 0 {
		ans = a/b + 1
	} else {
		ans = a / b
	}

	fmt.Fprintln(out, ans)
}
