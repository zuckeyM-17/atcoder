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

	var b string
	fmt.Fscan(in, &b)

	m := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}

	fmt.Fprintln(out, m[b])
}
