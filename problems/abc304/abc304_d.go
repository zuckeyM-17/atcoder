package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type tuple struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var w, h, n int
	fmt.Fscan(in, &w, &h)
	fmt.Fscan(in, &n)

	p, q := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}

	var a, b int

	fmt.Fscan(in, &a)
	as := make([]int, a+2)
	for i := 1; i <= a; i++ {
		fmt.Fscan(in, &as[i])
	}
	as[a+1] = w

	fmt.Fscan(in, &b)
	bs := make([]int, b+2)
	for i := 1; i <= b; i++ {
		fmt.Fscan(in, &bs[i])
	}
	bs[b+1] = h

	m := make(map[tuple]int)
	for i := 0; i < n; i++ {
		xIdx := sort.Search(len(as), func(k int) bool {
			return p[i] < as[k]
		})

		yIdx := sort.Search(len(bs), func(k int) bool {
			return q[i] < bs[k]
		})
		m[tuple{xIdx, yIdx}]++
	}

	var length, max int
	min := math.MaxInt64

	for _, v := range m {
		length++
		if max < v {
			max = v
		}
		if v < min {
			min = v
		}
	}
	if length < (a+1)*(b+1) {
		min = 0
	}

	fmt.Fprintln(out, min, max)
}
