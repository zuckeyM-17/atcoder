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

	var s string
	fmt.Fscan(in, &s)

	m := map[string]int{
		"tourist":    3858,
		"ksun48":     3679,
		"Benq":       3658,
		"Um_nik":     3648,
		"apiad":      3638,
		"Stonefeang": 3630,
		"ecnerwala":  3613,
		"mnbvmar":    3555,
		"newbiedmy":  3516,
		"semiexp":    3481,
	}

	fmt.Fprintln(out, m[s])
}
