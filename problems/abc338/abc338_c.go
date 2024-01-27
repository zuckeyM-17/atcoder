package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxMeals(N int, Q, A, B []int) int {
	maxPeopleA := getMaxPeople(N, Q, A)
	// maxPeopleB := getMaxPeople(N, Q, B)
	maxTotal := 0

	for peopleA := 0; peopleA <= maxPeopleA; peopleA++ {
		remainingQ := make([]int, N)
		for i := 0; i < N; i++ {
			remainingQ[i] = Q[i] - A[i]*peopleA
		}
		peopleB := getMaxPeople(N, remainingQ, B)
		total := peopleA + peopleB
		if total > maxTotal {
			maxTotal = total
		}
	}

	return maxTotal
}

func getMaxPeople(N int, Q, R []int) int {
	maxPeople := int(^uint(0) >> 1)
	for i := 0; i < N; i++ {
		if R[i] > 0 {
			people := Q[i] / R[i]
			if people < maxPeople {
				maxPeople = people
			}
		}
	}
	return maxPeople
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	q, a, b := make([]int, n), make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	fmt.Fprintln(out, maxMeals(n, q, a, b))
}
