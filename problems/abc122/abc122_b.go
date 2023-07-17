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

	var max int

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isACGT(s[i:j+1]) && max < j-i+1 {
				max = j - i + 1
			}
		}
	}

	fmt.Fprintln(out, max)
}

func isACGT(s string) bool {
	for _, r := range s {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return false
		}
	}
	return true
}
