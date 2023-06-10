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

	var p, q string
	fmt.Fscan(in, &p, &q)

	m := map[string]int{
		"A": 0,
		"B": 3,
		"C": 4,
		"D": 8,
		"E": 9,
		"F": 14,
		"G": 23,
	}

	x := abs(m[p] - m[q])
	fmt.Fprintln(out, x)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
