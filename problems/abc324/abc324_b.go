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

	result := isPowerOf2And3(n)

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPowerOf2And3(n int) bool {
	for n > 0 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else {
			break
		}
	}
	return n == 1
}
