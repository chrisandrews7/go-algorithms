package main

import (
	"os"
	"reflect"
	"testing"
)

var graph WeightedUndirectedGraph

func TestMain(m *testing.M) {
	graph = NewWeightedUndirectedGraph()

	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")

	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 1)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("C", "E", 1)
	graph.AddEdge("D", "E", 1)
	graph.AddEdge("D", "F", 1)
	graph.AddEdge("E", "F", 1)

	os.Exit(m.Run())
}

func TestDepthFirstSearch(t *testing.T) {
	//          A
	//        /   \
	//       B     C
	//       |     |
	//       D --- E
	//        \   /
	//          F

	result := DepthFirstSearch(graph, "A")
	expectedResult := []VertexName{"A", "C", "E", "F", "D", "B"}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Traversal %s is in the wrong order", result)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	//          A
	//        /   \
	//       B     C
	//       |     |
	//       D --- E
	//        \   /
	//          F

	result := BreadthFirstSearch(graph, "A")
	expectedResult := []VertexName{"A", "B", "C", "D", "E", "F"}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Traversal %s is in the wrong order", result)
	}
}
