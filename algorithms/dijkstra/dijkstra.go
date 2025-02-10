func dijkstraDirected(times [][]int, n int, k int) map[int]int {
	graph := combineGraph(times)

	dist := make(map[int]int, n)
	for i := range n {
		dist[i+1] = math.MaxInt64
	}

	minHeap := &MinHeap{Item{k, 0}}

	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(Item)

		if item.Cost < dist[item.Node] {
			dist[item.Node] = item.Cost

			for _, neighbor := range graph[item.Node] {
				newCost := neighbor.Weight + item.Cost

				if newCost < dist[neighbor.To] {
					heap.Push(minHeap, Item{neighbor.To, newCost})
				}
			}
		}
	}

	return dist
}

type Edge struct {
	To     int
	Weight int
}

func combineGraph(edges [][]int) map[int][]Edge {
	graph := make(map[int][]Edge, len(edges)+1)

	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		graph[from] = append(graph[from], Edge{to, weight})
	}

	return graph
}

type Item struct {
	Node int
	Cost int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}