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

	fmt.Fprintln(out)
}

// xz≡y(mod998244353) を満たすような 0 以上 998244352 以下の整数 z を求める
func mod998244353(x, y int) int {
	return x * 998244353 / y
}
