package main

type VertexName string

type Vertex struct {
	name  VertexName
	edges []*Edge
}

type Edge struct {
	from   VertexName
	to     VertexName
	weight int
}

type WeightedUndirectedGraph struct {
	adjacencyList map[VertexName]*Vertex
}

func (graph *WeightedUndirectedGraph) AddVertex(name VertexName) {
	vertex := &Vertex{
		name: name,
	}

	graph.adjacencyList[name] = vertex
}

func (graph *WeightedUndirectedGraph) AddEdge(from VertexName, to VertexName, weight int) {
	edgeFrom := &Edge{
		from:   from,
		to:     to,
		weight: weight,
	}
	edgeTo := &Edge{
		from:   to,
		to:     from,
		weight: weight,
	}

	graph.adjacencyList[from].edges = append(graph.adjacencyList[from].edges, edgeFrom)
	graph.adjacencyList[to].edges = append(graph.adjacencyList[to].edges, edgeTo)
}

func NewWeightedUndirectedGraph() WeightedUndirectedGraph {
	return WeightedUndirectedGraph{
		adjacencyList: make(map[VertexName]*Vertex),
	}
}
