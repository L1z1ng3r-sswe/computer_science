package topological_sort

import "fmt"

func TopologicalSort(numVertices int, edges [][]int) ([]int, error) {
	adjList := make([][]int, numVertices)
	used := make([]int, numVertices)

	for _, edge := range edges {
		node1 := edge[0]
		node2 := edge[1]
		adjList[node1] = append(adjList[node1], node2)
		used[node2]++
	}

	queue := []int{}
	for node, usages := range used {
		if usages == 0 {
			queue = append(queue, node)
		}
	}

	var res []int
	var counter int
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		res = append(res, curr)
		counter++

		for _, availableCourse := range adjList[curr] {
			used[availableCourse]--
			if used[availableCourse] == 0 {
				queue = append(queue, availableCourse)
			}
		}
	}

	if counter != numVertices {
		return nil, fmt.Errorf("Cycle detected")
	}

	return res, nil
}
