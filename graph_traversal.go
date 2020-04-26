package main

type traversalStorage interface {
	Enqueue(value interface{})
	Dequeue() interface{}
	Len() int
}

func traverse(storage traversalStorage, graph WeightedUndirectedGraph, start VertexName) (results []VertexName) {
	visited := make(map[VertexName]bool)
	store := storage

	startName := graph.adjacencyList[start].name
	store.Enqueue(startName)
	visited[startName] = true

	for store.Len() > 0 {
		name := store.Dequeue()
		vertex := graph.adjacencyList[name.(VertexName)]

		results = append(results, vertex.name)

		for _, edge := range vertex.edges {
			if visited[edge.to] == false {
				visited[edge.to] = true
				store.Enqueue(edge.to)
			}
		}
	}

	return
}

func DepthFirstSearch(graph WeightedUndirectedGraph, start VertexName) []VertexName {
	stack := NewStack()
	return traverse(&stack, graph, start)
}

func BreadthFirstSearch(graph WeightedUndirectedGraph, start VertexName) []VertexName {
	queue := NewQueue()
	return traverse(&queue, graph, start)
}
