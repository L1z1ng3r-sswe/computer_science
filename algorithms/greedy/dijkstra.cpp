#include <iostream>
#include <vector>
#include <unordered_map>
#include <limits.h>
#include <queue>

struct Edge {
  int To;
  int Weight;
};

std::vector<std::vector<Edge>> combineGraph(std::vector<std::vector<int>> edges, int n) {
  std::vector<std::vector<Edge>> graph(n);

  for (const std::vector<int>& edge : edges) {
    int from = edge[0], to = edge[1], weight = edge[2];

    graph[from].push_back(Edge{to, weight});
  }

  return graph;
}

struct Node {
  int node;
  int dist;

  bool operator>(const Node& node2) const {
    return dist > node2.dist;
  }
};

std::unordered_map<int, int> dijkstra(std::vector<std::vector<int>> edges, int n, int source) {
  std::vector<std::vector<Edge>> graph = combineGraph(edges, n);

  std::unordered_map<int, int> dist (n);
  for (int i = 0; i < n; i++) {
    dist[i] = INT_MAX;
  }
  dist[source] = 0;

  std::priority_queue<Node, std::vector<Node>, std::greater<Node>> minHeap;
  minHeap.push(Node{source, 0});

  while (!minHeap.empty()) {
    Node curr = minHeap.top();
    minHeap.pop();

    int currNode = curr.node;
    int currDist = curr.dist;

    if (currDist <= dist[currNode]) {
      for (const Edge& neighbor : graph[currNode]) {
        int newDist = currDist + neighbor.Weight;

        if (newDist < dist[neighbor.To]) {
          dist[neighbor.To] = newDist;
          minHeap.push(Node{neighbor.To, newDist});
        }
      }
    }
  }

  return dist;
}