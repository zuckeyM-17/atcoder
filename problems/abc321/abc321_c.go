package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)

	digits := "0123456789"

	if k <= 9 {
		fmt.Fprintln(out, k)
		return
	} else {
		k -= 9
		var prev []string
		for i := 0; i < 10; i++ {
			prev = append(prev, strconv.Itoa(i))
		}
		for {
			var tmp []string
			for _, next := range digits {
				for _, e := range prev {
					if int(next) > int(e[0]) {
						tmp = append(tmp, string(next)+e)
					}
				}
			}
			prev = tmp
			if k <= len(prev) {
				break
			} else {
				k -= len(prev)
			}
		}
		fmt.Fprintln(out, prev[k-1])
	}
}
