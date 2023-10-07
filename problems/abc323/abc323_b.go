package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type person struct {
	score  int
	number int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	persons := make([]person, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		persons[i] = person{0, i}

		for j := 0; j < len(s); j++ {
			if s[j] == 'o' {
				persons[i].score++
			}
		}
	}

	sort.Slice(persons, func(i, j int) bool {
		if persons[i].score == persons[j].score {
			return persons[i].number < persons[j].number
		} else {
			return persons[i].score > persons[j].score
		}
	})

	for i, v := range persons {
		fmt.Fprint(out, v.number+1)
		if i != n-1 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out)
		}
	}
}
