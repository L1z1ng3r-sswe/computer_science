package dijkstra

import (
	"container/heap"
	"math"
)

func dijkstra(edges [][]int, n int, k int) int {
	graph := buildGraph(n, edges)

	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt64
	}
	distances[k-1] = 0

	minHeap := &MinHeap{&Item{k, 0}}

	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(*Item)

		for _, edge := range graph[item.vertex-1] {
			newCost := item.distance + edge.weight

			if newCost < distances[edge.target-1] {
				distances[edge.target-1] = newCost
				heap.Push(minHeap, &Item{edge.target, newCost})
			}
		}
	}

	res := -1
	for _, dist := range distances {
		if dist == math.MaxInt64 {
			return -1
		}
		if dist > res {
			res = dist
		}
	}

	return res
}

func buildGraph(n int, edges [][]int) [][]*Edge {
	graph := make([][]*Edge, n)

	for _, edge := range edges {
		source, target, weight := edge[0], edge[1], edge[2]
		graph[source-1] = append(graph[source-1], &Edge{target, weight})
	}

	return graph
}

type Edge struct {
	target, weight int
}

type MinHeap []*Item

type Item struct {
	vertex, distance int
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h *MinHeap) Push(i interface{}) {
	*h = append(*h, i.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}
