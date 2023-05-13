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

	var s, t string
	fmt.Fscan(in, &s)
	fmt.Fscan(in, &t)

	ms, mt := make(map[rune]int), make(map[rune]int)
	for i := 0; i < len(s); i++ {
		ms[rune(s[i])]++
		mt[rune(t[i])]++
	}

	diffST, diffTS := make(map[rune]int), make(map[rune]int)
	ats, att := 0, 0
	for k, v := range ms {
		if '@' == k {
			ats = v
		} else if _, ok := mt[k]; !ok {
			diffST[k] = v
		} else if v != mt[k] {
			diffST[k] = v - mt[k]
		}
	}

	for k, v := range mt {
		if '@' == k {
			att = v
		} else if _, ok := ms[k]; !ok {
			diffTS[k] = v
		} else if v != ms[k] {
			diffTS[k] = v - ms[k]
		}
	}

	// fmt.Fprintln(out, diffST, diffTS, ats, att)

	uncoms := []rune{'b', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 's', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, k := range uncoms {
		if _, ok := diffST[k]; ok {
			fmt.Fprintln(out, "No")
			return
		} else if _, ok := diffTS[k]; ok {
			fmt.Fprintln(out, "No")
			return
		}
	}

	// fmt.Fprintln(out, diffST, diffTS, ats, att)

	coms := []rune{'a', 'c', 'd', 'e', 'o', 'r', 't'}
	comss, comst := 0, 0
	for _, k := range coms {
		if _, ok := diffST[k]; ok {
			comss += diffST[k]
		}
		if _, ok := diffTS[k]; ok {
			comst += diffTS[k]
		}
	}

	// fmt.Fprintln(out, comss, comst, ats, att)
	if comss > att {
		fmt.Fprintln(out, "No")
		return
	}

	if comst > ats {
		fmt.Fprintln(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
}
