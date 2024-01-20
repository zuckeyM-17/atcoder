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

	res := map[int]bool{}
	nums := []int{1, 11, 111, 1111, 11111, 111111, 1111111, 11111111, 111111111, 1111111111, 11111111111, 111111111111}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				m := nums[i] + nums[j] + nums[k]
				if res[m] {
					continue
				} else {
					res[m] = true
				}
			}
		}
	}

	ints := keys(res)
	sort.Ints(ints)
	fmt.Fprintln(out, ints[n-1])
}

func keys(m map[int]bool) []int {
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}
