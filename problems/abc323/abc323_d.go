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

	slimes := map[int]int{}
	for i := 0; i < n; i++ {
		var s, c int
		fmt.Fscan(in, &s, &c)

		powers := recursiveMaxPowers(c)
		for _, v := range powers {
			addSlime(slimes, s<<v)
		}
	}

	fmt.Fprintln(out, len(slimes))
}

func addSlime(slimes map[int]int, s int) {
	if _, ok := slimes[s]; !ok {
		slimes[s] = 1
	} else {
		delete(slimes, s)
		addSlime(slimes, s*2)
	}
}

func recursiveMaxPowers(n int) []int {
	if n == 0 {
		return []int{}
	}
	power := maxPower(n)
	return append(recursiveMaxPowers(n-(1<<power)), power)
}

// 与えられた数字よりも小さい、最大の2のべき乗の指数を返す
func maxPower(n int) int {
	if n == 0 {
		return 0
	}
	power := 0
	for n > 1 {
		n /= 2
		power++
	}
	return power
}
