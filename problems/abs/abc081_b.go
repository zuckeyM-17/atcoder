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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	action := 0
	for {
		for i, v := range a {
			if v%2 > 0 {
				fmt.Fprintln(out, action)
				return
			} else {
				a[i] = v / 2
			}
		}
		action++
	}
}
