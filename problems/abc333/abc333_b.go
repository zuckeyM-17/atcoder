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
	var t string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &t)

	m := map[string]int{
		"AB": 1,
		"BA": 1,
		"BC": 1,
		"CB": 1,
		"CD": 1,
		"DC": 1,
		"DE": 1,
		"ED": 1,
		"EA": 1,
		"AE": 1,
		"AC": 2,
		"CA": 2,
		"AD": 2,
		"DA": 2,
		"BD": 2,
		"DB": 2,
		"BE": 2,
		"EB": 2,
		"CE": 2,
		"EC": 2,
	}
	if m[s] == m[t] {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
