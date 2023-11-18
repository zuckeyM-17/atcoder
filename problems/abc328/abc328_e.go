package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Edge struct {
	src, dest, weight int
}

type Graph struct {
	vertices int
	edges    []Edge
}

func newGraph(vertices int) *Graph {
	return &Graph{vertices, []Edge{}}
}

func (g *Graph) addEdge(src, dest, weight int) {
	g.edges = append(g.edges, Edge{src, dest, weight})
}

type UnionFind struct {
	parent []int
}

func newUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	uf.parent[uf.find(x)] = uf.find(y)
}

func mst(g *Graph) []*Edge {
	result := make([]*Edge, 0)
	i, e := 0, 0

	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	uf := newUnionFind(g.vertices)

	for e < g.vertices-1 {
		nextEdge := g.edges[i]
		i++

		x := uf.find(nextEdge.src)
		y := uf.find(nextEdge.dest)

		if x != y {
			result = append(result, nextEdge)
			uf.union(x, y)
			e++
		}
	}
	return result
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	g := newGraph(n)

	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		g.addEdge(u-1, v-1, w)
	}
	mst := mst(g)

	ans := 0
	for _, e := range mst {
		ans += e.weight
		ans = mod(ans, k)
	}

	fmt.Fprintln(out, ans)
}
