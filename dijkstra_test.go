package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := NewWeightedUndirectedGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")

	graph.AddEdge("A", "B", 3)
	graph.AddEdge("A", "C", 1)
	graph.AddEdge("B", "E", 1)
	graph.AddEdge("B", "C", 7)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("D", "C", 2)
	graph.AddEdge("E", "D", 7)

	// A -3- B -1- E
	// |    / \    |
	// 1   7    1  7
	// | /       \ |
	// C --- 2 --- D

	distance, route := Dijkstra(graph, "C", "E")
	expectedPath := []VertexName{"C", "D", "B", "E"}

	if distance != 4 || !reflect.DeepEqual(route, expectedPath) {
		t.Errorf("%s (%d) is not the shortest route", route, distance)
	}
}
