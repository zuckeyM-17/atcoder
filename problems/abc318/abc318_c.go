package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, d, p int
	fmt.Fscan(in, &n, &d, &p)

	f := make([]int, n)

	passFee := p / d

	type tuple struct {
		date, fee int
	}

	overDate := []tuple{}
	fee := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &f[i])
		if f[i] > passFee {
			overDate = append(overDate, tuple{date: i, fee: f[i]})
		} else {
			fee += f[i]
		}
	}

	if len(overDate) == 0 {
		ans := 0
		for i := range f {
			ans += f[i]
		}
		fmt.Fprintln(out, ans)
		return
	}

	sort.Slice(overDate, func(i, j int) bool {
		return overDate[i].fee > overDate[j].fee
	})

	passCount := len(overDate) / d

	rest := overDate[passCount*d:]
	restFee := 0
	for _, r := range rest {
		restFee += r.fee
	}
	if restFee > p {
		passCount++
	} else {
		fee += restFee
	}

	fmt.Fprintln(out, fee+passCount*p)
}
