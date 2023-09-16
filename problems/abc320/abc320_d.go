package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	x, y  int
	added bool
}

type info struct {
	a, b, x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	p := make([]person, n+1)
	p[1] = person{0, 0, true}

	var a, b, x, y int

	restInfo := []info{}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b, &x, &y)

		if !p[a].added && !p[b].added {
			restInfo = append(restInfo, info{a, b, x, y})
		} else if !p[a].added {
			p[a] = person{p[b].x - x, p[b].y - y, true}
		} else {
			p[b] = person{p[a].x + x, p[a].y + y, true}
		}
	}

	count := 0
	for len(restInfo) > 0 && count < 10 {
		for i := 0; i < len(restInfo); i++ {
			a, b, x, y = restInfo[i].a, restInfo[i].b, restInfo[i].x, restInfo[i].y

			if !p[a].added && !p[b].added {
				continue
			} else if !p[a].added {
				p[a] = person{p[b].x - x, p[b].y - y, true}
			} else {
				p[b] = person{p[a].x + x, p[a].y + y, true}
			}

			restInfo = append(restInfo[:i], restInfo[i+1:]...)
			i--
		}
		count++
	}

	for i := 1; i <= n; i++ {
		if p[i] == (person{}) {
			fmt.Fprintln(out, "undecidable")
		} else {
			fmt.Fprintln(out, p[i].x, p[i].y)
		}
	}
}
