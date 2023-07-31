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

	var x, y, a, b int
	fmt.Fscan(in, &x, &y, &a, &b)

	var ans int
	for {
		p := float64(x) * float64(a)
		q := float64(x) + float64(b)
		if p >= q {
			break
		}
		if p >= float64(y) {
			break
		}
		x *= a
		ans++
	}
	ans += (y - x - 1) / b
	fmt.Fprintln(out, ans)
}
