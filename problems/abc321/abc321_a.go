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

	var n string
	fmt.Fscan(in, &n)

	for i := 0; i < len(n)-1; i++ {
		if int(n[i]) <= int(n[i+1]) {
			fmt.Fprintln(out, "No")
			return
		}
	}

	fmt.Fprintln(out, "Yes")
}
