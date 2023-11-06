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

	m := map[int]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		m[a]++
	}

	ans := 0
	for i := 1; i <= 100; i++ {
		if m[i] >= 3 {
			ans += m[i] * (m[i] - 1) * (m[i] - 2) / 6
		}
	}

	fmt.Fprintln(out, ans)
}
