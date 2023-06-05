package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type box struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	boxes := make([]box, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		boxes[i] = box{x, y}
	}

	sort.Slice(boxes, func(i, j int) bool {
		if boxes[i].x == boxes[j].x {
			return boxes[i].y > boxes[j].y
		}
		return boxes[i].x < boxes[j].x
	})

	l := []int{}
	for _, box := range boxes {
		pos := sort.SearchInts(l, box.y)
		if pos == len(l) {
			l = append(l, box.y)
		} else {
			l[pos] = box.y
		}
	}

	fmt.Fprintln(out, len(l))
}
