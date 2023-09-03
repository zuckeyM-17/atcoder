package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"os"
)

type Edge struct {
	i, j, w int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, w int
	fmt.Fscan(in, &n)

	edges := make([][]Edge, n-1)
	for i := range edges {
		edges[i] = make([]edge, n-1-i)
		for j := range edges[i] {
			fmt.Fscan(in, &w)
			edges[i][j] = Edge{i: i, j: j + i + 1, w: w}
		}
	}


	for 

	fmt.Fprintln(out)
}
