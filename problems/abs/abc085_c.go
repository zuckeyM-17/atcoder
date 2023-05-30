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

	var n, y int
	fmt.Fscan(in, &n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			for k := 0; k <= n-i-j; k++ {
				if i+j+k == n && i*10000+j*5000+k*1000 == y {
					fmt.Fprintln(out, i, j, k)
					return
				}
			}
		}
	}
	fmt.Fprintln(out, -1, -1, -1)
}
