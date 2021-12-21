package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

type Node struct {
	path     *Vertex
	value    *Vertex
	priority float64
}

type PQ []*Node

func (pq *PQ) enqueue(vertex *Vertex, priority float64, path *Vertex) {
	newNode := Node{value: vertex, priority: priority, path: path}
	*pq = append(*pq, &newNode)
	index := len(*pq) - 1
	current := (*pq)[index]

	for index > 0 {
		parentIndex := ((index - 1) / 2)
		parent := (*pq)[parentIndex]

		if parent.priority >= current.priority {
			(*pq)[parentIndex], (*pq)[index] = current, parent
			index = parentIndex
		} else {
			break
		}
	}
}

func (pq *PQ) dequeue() *Node {
	if len(*pq) == 0 {
		return nil
	}

	max := (*pq)[0]
	end := (*pq)[len((*pq))-1]
	(*pq) = (*pq)[:len((*pq))-1]
	if len(*pq) == 0 {
		return max
	}

	(*pq)[0] = end

	index := 0
	length := len(*pq)
	current := (*pq)[0]
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		var leftChild *Node
		var rightChild *Node
		swap := -1

		if leftChildIndex < length {
			leftChild = (*pq)[leftChildIndex]
			if leftChild.priority < current.priority {
				swap = leftChildIndex
			}
		}
		if rightChildIndex < length {
			rightChild = (*pq)[rightChildIndex]
			if (swap == -1 && rightChild.priority < current.priority) || (swap != -1 && rightChild.priority < leftChild.priority) {
				swap = rightChildIndex
			}
		}

		if swap == -1 {
			break
		}
		(*pq)[index], (*pq)[swap] = (*pq)[swap], current
		index = swap
	}

	return max
}

type Vertex struct {
	Key      string
	Cost     int
	Vertices map[*Vertex]float64
}

type Graph struct {
	Vertices map[string]*Vertex
}

func (g *Graph) AddVertex(key string, cost string) *Vertex {
	i, _ := strconv.Atoi(cost)
	v := &Vertex{Key: key, Cost: i, Vertices: make(map[*Vertex]float64)}
	g.Vertices[key] = v
	return v
}

func (v *Vertex) AddEdge(vertex *Vertex, cost int) {
	v.Vertices[vertex] = float64(cost)
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func GraphFromInput(lines []string, numberOfGrids int) (*Graph, *Vertex, *Vertex) {
	graph := NewGraph()
	width := len(lines[0])
	height := len(lines)

	nodeSlice := make([][]*Vertex, height*numberOfGrids)

	for y, line := range lines {
		if len(nodeSlice[y]) <= 0 {
			nodeSlice[y] = make([]*Vertex, height*numberOfGrids)
		}
		for x, char := range line {
			v := graph.AddVertex(fmt.Sprintf("%d-%d", y, x), string(char))
			nodeSlice[y][x] = v
		}
	}

	for dy := 0; dy < numberOfGrids; dy++ {
		for dx := 0; dx < numberOfGrids; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					value := nodeSlice[y][x].Cost + dx + dy
					if value > 9 {
						value = (value % 9)
					}
					yKey := y + (dy * (height))
					xKey := x + (dx * (width))
					if len(nodeSlice[yKey]) <= 0 {
						nodeSlice[yKey] = make([]*Vertex, width*numberOfGrids)
					}
					v := graph.AddVertex(fmt.Sprintf("%d-%d", yKey, xKey), fmt.Sprint(value))
					nodeSlice[yKey][xKey] = v
				}
			}
		}
	}

	for nodeLineIndex, line := range nodeSlice {
		for nodeIndex, node := range line {
			if nodeIndex+1 <= len(line)-1 {
				node.AddEdge(line[nodeIndex+1], line[nodeIndex+1].Cost)
			}
			if nodeIndex-1 >= 0 {
				node.AddEdge(line[nodeIndex-1], line[nodeIndex-1].Cost)
			}
			if nodeLineIndex+1 <= len(nodeSlice)-1 {
				node.AddEdge(nodeSlice[nodeLineIndex+1][nodeIndex], nodeSlice[nodeLineIndex+1][nodeIndex].Cost)
			}
			if nodeLineIndex-1 >= 0 {
				node.AddEdge(nodeSlice[nodeLineIndex-1][nodeIndex], nodeSlice[nodeLineIndex-1][nodeIndex].Cost)
			}
		}
	}

	return graph, nodeSlice[0][0], nodeSlice[len(nodeSlice)-1][len(nodeSlice[0])-1]
}

type CostAndVertex struct {
	cost   float64
	vertex *Vertex
}

func dijkstra(startingVertex *Vertex) map[*Vertex]CostAndVertex {
	routesFromVertex := map[*Vertex]CostAndVertex{}
	routesFromVertex[startingVertex] = CostAndVertex{cost: 0, vertex: startingVertex}

	pqueue := PQ{}

	for k, v := range startingVertex.Vertices {
		pqueue.enqueue(k, v, startingVertex)
	}

	for {
		nextV := pqueue.dequeue()
		if nextV == nil {
			break
		}

		if _, ok := routesFromVertex[nextV.value]; ok {
			continue
		}
		var cost float64
		if _, ok := routesFromVertex[nextV.path]; ok {
			cost = routesFromVertex[nextV.path].cost + float64(nextV.value.Cost)
		} else {
			cost = float64(nextV.value.Cost)
		}

		routesFromVertex[nextV.value] = CostAndVertex{cost: cost, vertex: nextV.path}

		for k, v := range nextV.value.Vertices {
			pqueue.enqueue(k, cost+v, nextV.value)
		}
	}
	return routesFromVertex
}

func part1(lines []string) int {
	_, start, end := GraphFromInput(lines, 1)

	h := dijkstra(start)
	return int(h[end].cost)
}

func part2(lines []string) int {
	_, start, end := GraphFromInput(lines, 5)

	h := dijkstra(start)
	return int(h[end].cost)
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
