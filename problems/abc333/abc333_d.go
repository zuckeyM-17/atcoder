package main

import (
	"bufio"
	"fmt"
	"os"
)

func minOperations(N int, edges [][]int) int {
	if N == 1 {
		return 1
	}

	// 隣接リストと度数の配列を作成
	adjList := make([][]int, N+1)
	degree := make([]int, N+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
		degree[u]++
		degree[v]++
	}

	if degree[1] == 1 {
		return 1
	}

	// 葉のリストを初期化
	leaves := make([]int, 0)
	for i := 2; i <= N; i++ { // 頂点1は除外
		if degree[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	sumd := make([]int, N+1)
	for degree[1] > 1 {
		newLeaves := make([]int, 0)
		for _, leaf := range leaves {
			degree[leaf] = 0
			for _, neighbor := range adjList[leaf] {
				if degree[neighbor] > 0 {
					degree[neighbor]--
					sumd[neighbor] += sumd[leaf] + 1
					if degree[neighbor] == 1 && neighbor != 1 {
						newLeaves = append(newLeaves, neighbor)
					}
				}
			}
		}
		leaves = newLeaves
	}

	res := []int{}
	for _, v := range adjList[1] {
		res = append(res, sumd[v])
	}

	fmt.Println(adjList[1])
	fmt.Println(sumd)
	fmt.Println(res)

	return min(res) + 2
}

func min(l []int) int {
	res := l[0]
	for _, v := range l {
		if v < res {
			res = v
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		edges[i] = []int{u, v}
	}

	fmt.Fprintln(out, minOperations(n, edges))
}
