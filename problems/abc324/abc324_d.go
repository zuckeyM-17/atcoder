package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var s string
	fmt.Fscan(in, &s)

	if n == 1 {
		if s == "1" || s == "4" || s == "9" || s == "0" {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
		return
	}

	squares := listSquares(n)
	ans := 0

	s = sortStringNumbers(s)

	for _, square := range squares {
		str := strconv.Itoa(square)
		numberOfZeros := len(s) - len(str)
		str = strings.Repeat("0", numberOfZeros) + str

		if sortStringNumbers(str) == s {
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}

func sortStringNumbers(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func listSquares(n int) []int {
	minNPlus1 := int(math.Pow10(n))
	endRoot := int(math.Sqrt(float64(minNPlus1)))

	squares := make([]int, endRoot+1)
	for i := 1; i < endRoot; i++ {
		squares[i] = i * i
	}
	return squares
}
