package main

type vertexName string

type vertex struct {
	name  vertexName
	edges []*edge
}

type edge struct {
	from   vertexName
	to     vertexName
	weight int
}

type WeightedDirectedGraph struct {
	adjacencyList map[vertexName]*vertex
}

func (graph *WeightedDirectedGraph) AddVertex(name vertexName) {
	vertex := vertex{
		name: name,
	}

	graph.adjacencyList[name] = &vertex
}

func (graph *WeightedDirectedGraph) AddEdge(from vertexName, to vertexName, weight int) {
	edge := edge{
		from:   from,
		to:     to,
		weight: weight,
	}

	graph.adjacencyList[from].edges = append(graph.adjacencyList[from].edges, &edge)
}

func NewWeightedDirectedGraph() WeightedDirectedGraph {
	return WeightedDirectedGraph{
		adjacencyList: make(map[vertexName]*vertex),
	}
}
