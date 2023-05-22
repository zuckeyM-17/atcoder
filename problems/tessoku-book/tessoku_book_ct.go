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

	var n int
	fmt.Fscan(in, &n)

	var str string
	fmt.Fscan(in, &str)

	s := []rune(str)

	fmt.Fprintln(out, s)
}
