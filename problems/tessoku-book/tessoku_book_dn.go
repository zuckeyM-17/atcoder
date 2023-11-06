package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, y int
	fmt.Fscan(in, &x, &y)

	cnt := 0
	pairs := []pair{{x, y}}

	for x > 1 || y > 1 {
		if x > y {
			x -= y
		} else {
			y -= x
		}
		cnt++
		pairs = append(pairs, pair{x, y})
	}

	fmt.Fprintln(out, cnt)
	for i := 1; i < len(pairs); i++ {
		fmt.Fprintln(out, pairs[cnt-i].x, pairs[cnt-i].y)
	}

}
