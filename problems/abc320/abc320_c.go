package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func str2map(s string) map[string][]int {
	m := map[string][]int{}
	for i, r := range s {
		str := string(r)
		if _, ok := m[str]; !ok {
			m[str] = []int{i}
		} else {
			m[str] = append(m[str], i)
		}
	}
	return m
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &m)

	var s1, s2, s3 string
	fmt.Fscan(in, &s1, &s2, &s3)

	m1 := str2map(s1)
	m2 := str2map(s2)
	m3 := str2map(s3)

	candidates := []string{}

	for i := 0; i <= 9; i++ {
		stri := strconv.Itoa(i)
		if _, ok := m1[stri]; !ok {
			continue
		}

		if _, ok := m2[stri]; !ok {
			continue
		}

		if _, ok := m3[stri]; !ok {
			continue
		}

		candidates = append(candidates, stri)
	}

	if len(candidates) == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	ans := 1000000000

	str1 := strings.Repeat(s1, 3)
	str2 := strings.Repeat(s2, 3)
	str3 := strings.Repeat(s3, 3)

	slots := [][]string{
		{str1, str2, str3},
		{str1, str3, str2},
		{str2, str1, str3},
		{str2, str3, str1},
		{str3, str1, str2},
		{str3, str2, str1},
	}

	for _, c := range candidates {
		for _, slot := range slots {
			cur := 0
			for t := 0; t < m*3; t++ {
				if string(slot[cur][t]) == c {
					cur++
					if cur == 3 {
						ans = min(ans, t)
						break
					}
				}
			}
		}
	}

	fmt.Fprintln(out, ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
