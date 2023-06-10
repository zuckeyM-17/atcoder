package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	noneStrRow, noneStrCol := strings.Repeat(".", w), strings.Repeat(".", h)
	cookieRows, cookieColumns := []int{}, []int{}
	s := make([][]string, h+1)
	for i := 1; i <= h; i++ {
		s[i] = make([]string, w+1)
		var str string
		fmt.Fscan(in, &str)

		if str != noneStrRow {
			cookieRows = append(cookieRows, i)
		}
		for j := 1; j <= w; j++ {
			s[i][j] = string(str[j-1])
		}
	}

	for j := 1; j <= w; j++ {
		str := ""
		for i := 1; i <= h; i++ {
			str += s[i][j]
		}
		if str != noneStrCol {
			cookieColumns = append(cookieColumns, j)
		}
	}

	for i := cookieRows[0]; i <= cookieRows[len(cookieRows)-1]; i++ {
		for j := cookieColumns[0]; j <= cookieColumns[len(cookieColumns)-1]; j++ {
			if s[i][j] == "." {
				fmt.Fprintln(out, i, j)
			}
		}
	}
}
