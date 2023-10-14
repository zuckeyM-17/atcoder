package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var t string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &t)

	var s string
	var ans []string
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s)

		if compare(s, t) {
			ans = append(ans, strconv.Itoa(i))
		}
	}

	fmt.Fprintln(out, len(ans))
	fmt.Fprintln(out, strings.Join(ans, " "))
}

func compare(s, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if abs(lenS-lenT) > 1 {
		return false
	}

	if lenS == lenT {
		if s == t {
			return true
		} else {
			diffCount := 0
			for i := 0; i < lenS; i++ {
				if s[i] != t[i] {
					diffCount++
				}
				if diffCount > 1 {
					return false
				}
			}
			if diffCount <= 1 {
				return true
			}
		}
	}

	if lenS > lenT {
		s, t = t, s
	}

	i, j := 0, 0
	diffCount := 0

	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			diffCount++
			j++ // tの方を1文字進める
			if diffCount > 1 {
				return false
			}
			continue
		}
		i++
		j++
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
