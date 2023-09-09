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

	var s11, s12, s13, s21, s22, s23, s31, s32, s33 int
	fmt.Fscan(in, &s11, &s12, &s13, &s21, &s22, &s23, &s31, &s32, &s33)

	fmt.Fprintln(out, s11, s21, s31, s12, s22, s32, s13, s23, s33)
}
