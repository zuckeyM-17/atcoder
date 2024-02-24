package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b, gcdValue int) int {
	return a / gcdValue * b
}

func findKthNumber(n, m, k int) int {
	gcdValue := gcd(n, m)
	lcmValue := lcm(n, m, gcdValue)

	if lcmValue == n {
		return m + n*(k-1)
	}

	if lcmValue == m {
		return n + m*(k-1)
	}

	countPerPeriod := lcmValue/n + lcmValue/m - 2

	a := k / countPerPeriod
	b := k % countPerPeriod

	if b != 0 {
		nn, mm, cur := 1, 1, 0
		for i := 1; i <= b; i++ {
			if nn*n < mm*m {
				cur = nn * n
				nn++
			} else {
				cur = mm * m
				mm++
			}
		}
		return a*lcmValue + cur
	} else {
		nn, mm, cur := 1, 1, 0
		for i := 1; i <= countPerPeriod; i++ {
			if nn*n < mm*m {
				cur = nn * n
				nn++
			} else {
				cur = mm * m
				mm++
			}
		}
		return (a-1)*lcmValue + cur
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	fmt.Fprintln(out, findKthNumber(n, m, k))
}
