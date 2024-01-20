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

	if n == 1 {
		fmt.Fprintln(out, "1")
	} else if n == 2 {
		fmt.Fprintln(out, "22")
	} else if n == 3 {
		fmt.Fprintln(out, "333")
	} else if n == 4 {
		fmt.Fprintln(out, "4444")
	} else if n == 5 {
		fmt.Fprintln(out, "55555")
	} else if n == 6 {
		fmt.Fprintln(out, "666666")
	} else if n == 7 {
		fmt.Fprintln(out, "7777777")
	} else if n == 8 {
		fmt.Fprintln(out, "88888888")
	} else {
		fmt.Fprintln(out, "999999999")
	}
}
