package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 島の数とツアールートを受け取り、最適な橋を封鎖したときの最短ツアー長を返す関数
func findMinimumTourLength(n int, tour []int) int {
	// ツアーの初期長を計算
	totalLength := len(tour) - 1

	// 最短距離を無限大で初期化
	minDistance := math.MaxInt32

	// 各橋を検討
	for i := 0; i < len(tour)-1; i++ {
		// 橋を渡ると仮定した距離
		distance := getDistance(n, tour[i], tour[i+1])

		// 最短距離の更新
		if distance < minDistance {
			minDistance = distance
		}
	}

	// 最適な橋を封鎖した場合のツアー長を返す
	return totalLength - minDistance
}

// 島Aから島Bへの最短距離を計算する関数
func getDistance(n, a, b int) int {
	// 距離を計算
	distance := math.Abs(float64(a - b))

	// 環状の距離を考慮
	if distance > float64(n)/2 {
		distance = float64(n) - distance
	}

	return int(distance)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	tour := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &tour[i])
	}

	minLength := findMinimumTourLength(n, tour)

	fmt.Fprintln(out, "最短ツアー長:", minLength)
}
