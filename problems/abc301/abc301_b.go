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

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	answer := []int{}
	answer = append(answer, a[0])
	for i := 0; i < n-1; i++ {
		answer = append(answer, diffArr(a[i], a[i+1])...)
		answer = append(answer, a[i+1])
	}

	for i, v := range answer {
		if i == len(answer)-1 {
			fmt.Fprintln(out, v)
		} else {
			fmt.Fprint(out, v, " ")
		}
	}
}

func diffArr(a, b int) []int {
	res := []int{}

	if a > b {
		for i := a - 1; i > b; i-- {
			res = append(res, i)
		}
	} else {
		for i := a + 1; i < b; i++ {
			res = append(res, i)
		}
	}

	return res
}
