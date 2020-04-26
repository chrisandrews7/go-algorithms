package main

type vertex struct {
	name       VertexName
	distance   int
	fromVertex VertexName
}

type queuedItem struct {
	name     VertexName
	priority int
}

type verticies map[VertexName]*vertex

func Dijkstra(graph WeightedUndirectedGraph, start VertexName, end VertexName) (distance int, fastestPath []VertexName) {
	compare := func(a interface{}, b interface{}) bool {
		return a.(queuedItem).priority > b.(queuedItem).priority
	}
	queue := NewPriorityQueue(compare)
	verticies := make(verticies)

	// Setup
	for _, v := range graph.adjacencyList {
		verticies[v.name] = &vertex{
			distance: 999999,
			name:     v.name,
		}
		if v.name == start {
			verticies[v.name].distance = 0
		}

		queue.Enqueue(queuedItem{
			name:     v.name,
			priority: verticies[v.name].distance,
		})
	}

	// Traverse routes
	for queue.Len() > 0 {
		name := queue.Dequeue().(queuedItem).name

		if name == end {
			break
		}

		for _, edge := range graph.adjacencyList[name].edges {
			distance := edge.weight + verticies[name].distance

			if verticies[edge.to].distance > distance {
				verticies[edge.to].distance = distance
				verticies[edge.to].fromVertex = edge.from

				queue.Enqueue(queuedItem{
					name:     edge.to,
					priority: distance,
				})
			}
		}
	}

	// Summarise
	previousVertex := verticies[end]
	for previousVertex != nil {
		fastestPath = append([]VertexName{previousVertex.name}, fastestPath...)
		previousVertex = verticies[previousVertex.fromVertex]
	}

	return verticies[end].distance, fastestPath
}
