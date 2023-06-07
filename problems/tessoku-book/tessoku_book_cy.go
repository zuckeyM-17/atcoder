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

	deleted := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if deleted[i] {
			continue
		}
		for j := i * 2; j <= n; j += i {
			deleted[j] = true
		}
	}

	for i := 2; i <= n; i++ {
		if !deleted[i] {
			fmt.Fprintln(out, i)
		}
	}
}
