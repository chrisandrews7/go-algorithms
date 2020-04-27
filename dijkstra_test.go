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
	graph.AddVertex("F")
	graph.AddVertex("G")

	graph.AddEdge("A", "B", 3)
	graph.AddEdge("A", "C", 1)
	graph.AddEdge("B", "E", 3)
	graph.AddEdge("B", "C", 7)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("D", "C", 2)
	graph.AddEdge("E", "D", 7)
	graph.AddEdge("E", "G", 2)
	graph.AddEdge("F", "G", 6)

	// A -3- B -3- E -2- G
	// |    / \    |     |
	// 1   7    1  7     6
	// | /       \ |     |
	// C --- 2 --- D -1- F

	distance, route := Dijkstra(graph, "C", "G")
	expectedPath := []VertexName{"C", "D", "B", "E", "G"}

	if distance != 8 || !reflect.DeepEqual(route, expectedPath) {
		t.Errorf("%s (%d) is not the shortest route", route, distance)
	}
}
