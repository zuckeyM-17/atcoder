package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	m := map[int]bool{}

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if m[a] {
			continue
		} else {
			m[a] = true
		}
	}
	is := getKeys(m)
	sort.Ints(is)

	fmt.Fprintln(out, is[len(is)-2])
}

func getKeys(m map[int]bool) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
