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

	var n, m int
	fmt.Fscan(in, &n, &m)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	permutations := permuteStrings(s)

	for _, p := range permutations {
		ok := true
		for i := 0; i < len(p)-1; i++ {
			if replacable(p[i], p[i+1]) {
				continue
			} else {
				ok = false
				break
			}
		}
		if ok {
			fmt.Fprintln(out, "Yes")
			return
		}
	}

	fmt.Fprintln(out, "No")
}

func replacable(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	for i, _ := range r1 {
		t := make([]rune, len(r1))
		copy(t, r1)
		t[i] = r2[i]
		if string(t) == string(r2) {
			return true
		}
	}

	return false
}

func permuteStrings(arr []string) [][]string {
	result := [][]string{}
	permute(arr, 0, &result)
	return result
}

func permute(arr []string, start int, result *[][]string) {
	if start >= len(arr)-1 {
		temp := make([]string, len(arr))
		copy(temp, arr)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[start], arr[i] = arr[i], arr[start]
		permute(arr, start+1, result)
		arr[start], arr[i] = arr[i], arr[start]
	}
}
