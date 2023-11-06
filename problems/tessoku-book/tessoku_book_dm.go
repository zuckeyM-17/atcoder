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

	m := make([]int, 100)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		m[a%100]++
	}

	ans := 0

	for i := 1; i < 50; i++ {
		if m[i] > 0 && m[100-i] > 0 {
			ans += m[i] * m[100-i]
		}
	}
	if m[0] > 1 {
		ans += m[0] * (m[0] - 1) / 2
	}
	if m[50] > 1 {
		ans += m[50] * (m[50] - 1) / 2
	}

	fmt.Fprintln(out, ans)
}
