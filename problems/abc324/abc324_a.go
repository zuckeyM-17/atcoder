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
	var num, a int
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Fscan(in, &num)
		} else {
			fmt.Fscan(in, &a)
			if num != a {
				fmt.Fprintln(out, "No")
				return
			}
		}
	}

	fmt.Fprintln(out, "Yes")
}
