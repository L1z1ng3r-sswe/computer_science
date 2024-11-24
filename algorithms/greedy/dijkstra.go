package dijkstra

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

func combineGraph(edges [][]int, n int) [][]Edge {
	var graph = make([][]Edge, n)

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], Edge{To: to, Weight: weight})
	}

	return graph
}

type Node struct {
	Node int
	Dist int
}

type MinHeap []*Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	lastVal := old[n-1]
	*h = old[:n-1]
	return lastVal
}

// len(edges) = n - 1
func dijkstraDirected(edges [][]int, n int, source int) map[int]int {
	graph := combineGraph(edges, n)

	dist := make(map[int]int, n)
	for node := range n {
		dist[node] = math.MaxInt64
	}
	dist[source] = 0

	minHeap := &MinHeap{}
	heap.Push(minHeap, &Node{Node: source, Dist: 0})

	for minHeap.Len() > 0 {
		curr := minHeap.Pop().(*Node)
		currNode := curr.Node
		currDist := curr.Dist

		if currDist <= dist[currNode] {
			for _, neighbor := range graph[currNode] {
				newDist := currDist + neighbor.Weight
				if newDist < dist[neighbor.To] {
					dist[neighbor.To] = newDist
					heap.Push(minHeap, &Node{Node: neighbor.To, Dist: newDist})
				}
			}
		}
	}

	return dist
}
