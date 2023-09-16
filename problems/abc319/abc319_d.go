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

	l := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}

	left := max(l...) - 1
	right := sum(l...) + len(l) - 1

	for {
		mid := (left + right) / 2
		if mid == left || mid == right {
			break
		}

		rows := 1
		length := 0

		for i, word := range l {
			length = length + word
			if i > 0 {
				length++
			}
			if length > mid {
				rows++
				length = word
			}
		}

		if rows > m {
			left = mid
		} else {
			right = mid
		}
	}

	fmt.Fprintln(out, right)
}

func max(arr ...int) int {
	m := arr[0]
	for _, v := range arr {
		if m < v {
			m = v
		}
	}
	return m
}

func sum(arr ...int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}
