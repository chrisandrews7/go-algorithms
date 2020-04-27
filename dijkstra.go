package main

type vertex struct {
	name       VertexName
	distance   int
	fromVertex VertexName
}

type verticies map[VertexName]*vertex

type queueItem struct {
	name     VertexName
	priority int
}

const MAX_INT = int(^uint(0) >> 1)

func Dijkstra(graph WeightedUndirectedGraph, start VertexName, end VertexName) (distance int, fastestPath []VertexName) {
	compare := func(a interface{}, b interface{}) bool {
		return a.(queueItem).priority > b.(queueItem).priority
	}
	queue := NewPriorityQueue(compare)
	verticies := make(verticies)

	// Setup
	for _, v := range graph.adjacencyList {
		verticies[v.name] = &vertex{
			distance: MAX_INT,
			name:     v.name,
		}
		if v.name == start {
			verticies[v.name].distance = 0
		}

		queue.Enqueue(queueItem{
			name:     v.name,
			priority: verticies[v.name].distance,
		})
	}

	// Traverse routes
	for queue.Len() > 0 {
		name := queue.Dequeue().(queueItem).name

		if name == end {
			break
		}

		for _, edge := range graph.adjacencyList[name].edges {
			distance := edge.weight + verticies[name].distance

			if verticies[edge.to].distance > distance {
				verticies[edge.to].distance = distance
				verticies[edge.to].fromVertex = edge.from

				queue.Enqueue(queueItem{
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
