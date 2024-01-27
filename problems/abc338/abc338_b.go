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

	m := map[rune]int{}

	for _, r := range s {
		m[r]++
	}

	var maxKey rune
	maxValue := 0
	for k, v := range m {
		if maxValue < v || (maxValue == v && k < maxKey) {
			maxKey = k
			maxValue = v
		}
	}

	fmt.Fprintln(out, string(maxKey))
}
