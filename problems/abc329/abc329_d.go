package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	n, v int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	max := person{0, 0}

	mp := make(map[int]person)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		if _, ok := mp[a]; !ok {
			mp[a] = person{a, 1}
		} else {
			mp[a] = person{a, mp[a].v + 1}
		}

		if max.n == 0 {
			max = mp[a]
			fmt.Fprintln(out, a)
			continue
		}

		if mp[a].v > max.v || (mp[a].v == max.v && mp[a].n < max.n) {
			fmt.Fprintln(out, mp[a].n)
			max = mp[a]
		} else {
			fmt.Fprintln(out, max.n)
		}
	}
}
