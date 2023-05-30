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

	words := []string{"dream", "dreamer", "erase", "eraser"}

	for len(s) > 0 {
		matched := false
		for _, w := range words {
			if strings.HasSuffix(s, w) {
				matched = true
				s = s[:len(s)-len(w)]
				break
			}
		}
		if !matched {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
