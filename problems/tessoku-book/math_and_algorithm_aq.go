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

	ans := 1
	cnt := 0
	for b > 1 {
		b /= 2
		cnt++
	}

	for i := 0; i < c; i++ {
		ans *= square
		ans %= 1000000007
	}
	if r > 0 {
		ans *= a
		ans %= 1000000007
	}

	fmt.Fprintln(out, ans)
}
