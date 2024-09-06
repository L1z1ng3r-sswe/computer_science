# Graph Data Structure

A graph is a collection of vertices (also called nodes) connected by edges. Vertices represent entities, and edges represent the connections or relationships between these entities.

### Parts of a Graph

- **Vertex (Node)**: A point in the graph representing an entity.
- **Edge**: A line connecting two vertices, representing the relationship between them.

![Graph Parts](./assets/basic_graph.png)

### Adjacency

- **Vertex Adjacency**: Two vertices are adjacent if they are connected by an edge.

![Vertex Adjacency](./assets/vertex_adjacency.png)

- **Edge Adjacency**: Two edges are adjacent if they share a common vertex (intermediary).

![Edge Adjacency](./assets/edge_adjacency.png)

### Degree

- **Degrie**: The number of adjacnent vertices (connected edges) of the vertex. In a directed graph, there are in-degrees (incoming edges) and out-degrees (outgoing edges).

![Degree](./assets/degree.png)

### Path

- **Path**: A sequence of vertices where each adjacent pair is connected by an edge.

![Path](./assets/path.png)

## Types of Graphs

### Trivial Graph

- A graph with only one vertex and no edges.  

![Trivial Graph](./assets/type_trivial.png)

### Null Graph

- A graph with no edges, but possibly multiple vertices.

![Null Graph](./assets/type_null.png)

### Connected Graph

- A graph where there is a path between every pair of vertices.

![Connected Graph](./assets/type_connected.png)

### Disconnected Graph

- A graph where at least one pair of vertices does not have a path between them.

![Disconnected Graph](./assets/type_disconnected.png)

### Cyclic Graph

- A graph that contains at least one cycle (a path that starts and ends at the same vertex).

![Cyclic Graph](./assets/type_cyclic.png)

### Acyclic Graph

- A graph with no cycles.

![Acyclic Graph](./assets/type_acyclic.png)

### Simple Graph

- A graph with no loops (edges connected at both ends to the same vertex) and no more than one edge between any two vertices.

![Simple Graph](./assets/type_simple.png)

### Multigraph

- A graph that may have multiple edges (parallel edges) between the same pair of vertices.

![Multi Graph](./assets/type_multigraph.png)

### Regular Graph 

- A graph where each vertex has the same number of edges.

![Regular Graph](./assets/type_regular.png)

### Complete Graph

- A graph where every pair of vertices is connected by an edge.

![Complete Graph](./assets/type_complete.png)

### Pseudo Graph

- A graph that allows loops and multiple edges between the same pair of vertices.

![Pseudo Graph](./assets/type_pseudo.png)

### Directed Graph (Digraph)

- A graph where edges have a direction, indicated by an arrow.

![Directed Graph](./assets/type_directed.png)

### Undirected Graph

- A graph where edges do not have a direction.

![Undirected Graph](./assets/type_undirected.png)

### Weighted Graph

- A graph where edges have weights or costs associated with them.

![Weighted Graph](./assets/type_weighted.png)

## Graph Representations

### Adjacency Matrix

- A 2D array where the cell at row i and column j represents the presence (and possibly the weight) of an edge between vertices i and j.

![Adjacency Matrix](./assets/adjacency_matrix.png)

### Adjacency List

- A collection of lists or arrays where each list represents a vertex and contains all the vertices adjacent to it.

![Adjacency List](./assets/adjacency_list.png)

