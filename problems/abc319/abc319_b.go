package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	js := make([]int, 0)
	for i := 1; i <= 9; i++ {
		if n%i == 0 {
			js = append(js, i)
		}
	}

	ans := ""
	for i := 0; i <= n; i++ {
		a := ""
		for _, j := range js {
			if i%(n/j) == 0 {
				a = strconv.Itoa(j)
				break
			}
		}
		if a == "" {
			ans += "-"
		} else {
			ans += a
		}
	}

	fmt.Fprintln(out, ans)
}
