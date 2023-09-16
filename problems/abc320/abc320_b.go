package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if len(max) >= j-i {
				break
			}
			if isPalindrome(s[i:j]) {
				max = s[i:j]
				break
			}
		}
	}
	return max
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[len(s)-1-i] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	fmt.Fprintln(out, len(longestPalindrome(s)))
}
