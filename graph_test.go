package main

import (
	"testing"
)

func TestWeightedUndirectedGraph(t *testing.T) {
	graph := NewWeightedUndirectedGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")

	graph.AddEdge("A", "D", 20)
	graph.AddEdge("A", "B", 5)
	graph.AddEdge("B", "C", 10)

	// A -20- D
	// |
	// 5
	// |
	// v
	// B -10- C

	if edge := graph.adjacencyList["A"].edges[0]; edge.to != "D" && edge.weight != 20 {
		t.Error("Expected A -20- D")
	}
	if edge := graph.adjacencyList["D"].edges[0]; edge.to != "A" && edge.weight != 20 {
		t.Error("Expected D -20- A")
	}
	if edge := graph.adjacencyList["A"].edges[1]; edge.to != "B" && edge.weight != 5 {
		t.Error("Expected A -5- B")
	}
	if edge := graph.adjacencyList["B"].edges[0]; edge.to != "A" && edge.weight != 5 {
		t.Error("Expected B -5- A")
	}
	if edge := graph.adjacencyList["B"].edges[1]; edge.to != "C" && edge.weight != 10 {
		t.Error("Expected B -10- C")
	}
	if edge := graph.adjacencyList["C"].edges[0]; edge.to != "B" && edge.weight != 10 {
		t.Error("Expected C -10- B")
	}
}
