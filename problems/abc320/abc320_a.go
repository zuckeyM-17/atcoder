package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b float64
	fmt.Fscan(in, &a, &b)

	fmt.Fprintln(out, int(math.Pow(a, b)+math.Pow(b, a)))
}
