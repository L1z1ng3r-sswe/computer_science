# Dijkstra's Algorithm

This documentation provides an overview of **Dijkstra's Algorithm** and its working principles. Implementations are available in both **C++** and **Go**, located in the same directory as `main.cpp` and `main.go`.

## Overview

Dijkstra's Algorithm is a graph search algorithm used to find the shortest path from a source node to all other nodes in a weighted graph. It guarantees the shortest path in graphs with non-negative edge weights.

### Key Concepts
- **Weighted Graph**: Nodes connected by edges with associated weights (distances or costs).
- **Shortest Path**: The minimum cost to traverse from the source to each node.
- **Priority Queue**: Efficiently selects the node with the smallest distance.

### Steps
1. **Initialization**:
   - Assign a distance of infinity to all nodes except the source, which starts with a distance of 0.
   - Use a priority queue to store nodes prioritized by their tentative distances.

2. **Relaxation**:
   - Pop the node with the smallest distance from the queue.
   - Update distances for its neighbors if a shorter path is found through the current node.

3. **Completion**:
   - Repeat until all nodes are visited or the shortest paths to the desired nodes are found.

## Visualization

For a step-by-step explanation of how Dijkstra's Algorithm works, refer to this video:  
[Dijkstra's Algorithm Visualization](https://www.youtube.com/watch?v=GazC3A4OQTE&ab_channel=ComputerScience).

The video demonstrates the algorithm's process, including the use of a priority queue and the relaxation of edges.

## Files

- **C++ Implementation**: `dijkstra.cpp`
- **Go Implementation**: `dijkstra.go`