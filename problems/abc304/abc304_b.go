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

	if 1000 > n {
		fmt.Fprintln(out, n)
	} else if 10000 > n {
		fmt.Fprintln(out, n/10*10)
	} else if 100000 > n {
		fmt.Fprintln(out, n/100*100)
	} else if 1000000 > n {
		fmt.Fprintln(out, n/1000*1000)
	} else if 10000000 > n {
		fmt.Fprintln(out, n/10000*10000)
	} else if 100000000 > n {
		fmt.Fprintln(out, n/100000*100000)
	} else {
		fmt.Fprintln(out, n/1000000*1000000)
	}
}
