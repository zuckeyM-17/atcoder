package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	start, end int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	pairs := make([]pair, n)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		 pairs[i] = pair{l, r}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].end < pairs[j].end
	})

	ans := 0
	now := 0
	for _, p := range pairs {
		if now <= p.start {
			ans++
			now = p.end
		}
	}

	fmt.Fprintln(out, ans)
}
