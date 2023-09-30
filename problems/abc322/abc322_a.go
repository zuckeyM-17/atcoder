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

	var str string
	fmt.Fscan(in, &str)

	ans := -1
	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			continue
		}

		if str[i-3:i] == "ABC" {
			ans = i - 2
			break
		}
	}

	fmt.Fprintln(out, ans)
}
