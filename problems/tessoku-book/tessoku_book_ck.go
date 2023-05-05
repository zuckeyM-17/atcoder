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

	var n float32
	fmt.Fscan(in, &n)

	l, r := float32(0), n
	mid := (l + r) / float32(2)

	for i := 0; i < 20; i++ {
		if round(mid*mid*mid+mid) == n {
			break
		}
		if round(mid*mid*mid+mid) < n {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}

	fmt.Fprintf(out, "%.6f\n", mid)
}

func round(f float32) float32 {
	return float32(int(f*1000)) / 1000
}
