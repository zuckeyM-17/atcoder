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

	target := "snuke"

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([]string, h+1)
	for i := 1; i <= h; i++ {
		fmt.Fscan(in, &s[i])

		if idx := strings.Index(s[i], target); idx != -1 {
			fmt.Fprintf(out, "%d %d\n", i, idx+1)
			fmt.Fprintf(out, "%d %d\n", i, idx+2)
			fmt.Fprintf(out, "%d %d\n", i, idx+3)
			fmt.Fprintf(out, "%d %d\n", i, idx+4)
			fmt.Fprintf(out, "%d %d\n", i, idx+5)
			return
		}

		if idx := strings.Index(Reverse(s[i]), target); idx != -1 {
			fmt.Fprintf(out, "%d %d\n", i, idx+1)
			fmt.Fprintf(out, "%d %d\n", i, idx)
			fmt.Fprintf(out, "%d %d\n", i, idx-1)
			fmt.Fprintf(out, "%d %d\n", i, idx-2)
			fmt.Fprintf(out, "%d %d\n", i, idx-3)
			return
		}
	}

	columns := make([]string, w+1)
	for i := 1; i <= w; i++ {
		var runes []rune
		for j := 1; j <= h; j++ {
			runes = append(runes, rune(s[j][i-1]))
		}
		columns[i] = string(runes)
	}

	for i := 1; i <= w; i++ {
		if idx := strings.Index(columns[i], target); idx != -1 {
			fmt.Fprintf(out, "%d %d\n", idx+1, i)
			fmt.Fprintf(out, "%d %d\n", idx+2, i)
			fmt.Fprintf(out, "%d %d\n", idx+3, i)
			fmt.Fprintf(out, "%d %d\n", idx+4, i)
			fmt.Fprintf(out, "%d %d\n", idx+5, i)
			return
		}

		if idx := strings.Index(Reverse(columns[i]), target); idx != -1 {
			fmt.Fprintf(out, "%d %d\n", h-idx, i)
			fmt.Fprintf(out, "%d %d\n", h-idx-1, i)
			fmt.Fprintf(out, "%d %d\n", h-idx-2, i)
			fmt.Fprintf(out, "%d %d\n", h-idx-3, i)
			fmt.Fprintf(out, "%d %d\n", h-idx-4, i)
			return
		}
	}

	fmt.Fprintln(out)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
