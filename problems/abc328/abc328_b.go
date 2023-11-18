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

	ans := 0
	for i := 1; i <= n; i++ {
		var d int
		fmt.Fscan(in, &d)
		if i < 10 {
			for j := 1; j <= d; j++ {
				if i == j {
					ans++
				} else if i == (j/10) && i == (j%10) {
					ans++
				}
			}
		}

		if i == 11 || i == 22 || i == 33 || i == 44 || i == 55 || i == 66 || i == 77 || i == 88 || i == 99 {
			t := i / 10
			for j := 1; j <= d; j++ {
				if t == j {
					ans++
				} else if t == (j/10) && t == (j%10) {
					ans++
				}
			}
		}
	}
	fmt.Fprintln(out, ans)
}
