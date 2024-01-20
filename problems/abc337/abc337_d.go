package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func DotCount(s string) int {
	count := 0
	for _, c := range s {
		if c == '.' {
			count++
		}
	}
	return count
}

func FindDotOStringsRegex(S string, k int) int {
	pattern := fmt.Sprintf("[.o]{%d}", k)
	re := regexp.MustCompile(pattern)

	min := k + 1
	for i := 0; i <= len(S)-k; i++ {
		match := re.FindString(S[i:])
		if match != "" {
			dotCount := DotCount(match)
			if min > dotCount {
				min = dotCount
			}
		}
	}

	return min
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w, k int
	fmt.Fscan(in, &h, &w, &k)

	s := make([][]rune, h)
	ans := k + 1
	for i := 0; i < h; i++ {
		s[i] = make([]rune, w)
		var line string
		fmt.Fscan(in, &line)

		dotCount := FindDotOStringsRegex(line, k)
		if ans > dotCount {
			ans = dotCount
		}

		for j, r := range line {
			s[i][j] = r
		}
	}

	for i := 0; i < w; i++ {
		var line string
		for j := 0; j < h; j++ {
			line += string(s[j][i])
		}

		dotCount := FindDotOStringsRegex(line, k)
		if ans > dotCount {
			ans = dotCount
		}
	}

	if ans == k+1 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, ans)
}
