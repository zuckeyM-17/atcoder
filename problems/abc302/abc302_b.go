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

	target := "snuke"

	var h, w int
	fmt.Fscan(in, &h, &w)

	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &s[i])
	}

	dh := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dw := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 8; k++ {
				str := ""
				for t := 0; t < 5; t++ {
					hi, wi := i+t*dh[k], j+t*dw[k]
					if hi < 0 || hi >= h || wi < 0 || wi >= w {
						break
					}
					str += string(s[hi][wi])
				}
				if str == target {
					for t := 0; t < 5; t++ {
						hi, wi := i+t*dh[k]+1, j+t*dw[k]+1
						fmt.Fprintln(out, hi, wi)
					}
					return
				}
			}
		}
	}
}
