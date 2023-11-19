package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	for i, c := range strings.Split(s, "") {
		if i < len(s)-1 {
			fmt.Fprint(out, c, " ")
		} else {
			fmt.Fprintln(out, c)
		}
	}
}
