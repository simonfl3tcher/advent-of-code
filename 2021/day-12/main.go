package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type Graph struct {
	Vertices map[string]*Vertex
}

func (g *Graph) AddVertex(key string) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
	if v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func NewGraphFromInput(lines []string) *Graph {
	graph := NewGraph()
	nodes := map[string][]string{}
	set := map[string]bool{}
	for _, line := range lines {
		n := strings.Split(line, "-")
		set[n[0]] = true
		set[n[1]] = true
		nodes[n[0]] = append(nodes[n[0]], n[1])
	}

	for k := range set {
		graph.AddVertex(k)
	}

	for k, v := range nodes {
		for _, e := range v {
			graph.AddEdge(k, e)
		}
	}
	return graph
}

type Vertex struct {
	Key      string
	Visited  bool
	Upper    bool
	Vertices map[string]*Vertex
}

func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:      key,
		Upper:    strings.ToUpper(key) == key,
		Vertices: map[string]*Vertex{},
	}
}

func findPaths(graph *Graph, start *Vertex, end *Vertex, current *Vertex, visited map[string]bool, singleSmallCaveVisited bool, startSeen bool) int {
	if current.Key == "start" && startSeen {
		return 0
	}

	if current.Key == "start" {
		startSeen = true
	}

	if current.Key == "end" {
		return 1
	}

	if !current.Upper && visited[current.Key] {
		if !singleSmallCaveVisited {
			singleSmallCaveVisited = true
		} else {
			return 0
		}
	}

	acc := 0
	for _, adjacent := range graph.Vertices[current.Key].Vertices {
		newMap := map[string]bool{}
		for k, v := range visited {
			newMap[k] = v
		}
		newMap[current.Key] = true

		acc += findPaths(graph, start, end, adjacent, newMap, singleSmallCaveVisited, startSeen)
	}

	return acc
}

func part1(graph *Graph) int {
	return findPaths(graph, graph.Vertices["start"], graph.Vertices["end"], graph.Vertices["start"], map[string]bool{}, true, false)
}

func part2(graph *Graph) int {
	return findPaths(graph, graph.Vertices["start"], graph.Vertices["end"], graph.Vertices["start"], map[string]bool{}, false, false)
}

func main() {
	lines := utils.FileLinesToSlice("input.txt")
	graph := NewGraphFromInput(lines)
	fmt.Printf("Part 1: %d\n", part1(graph))
	fmt.Printf("Part 2: %d\n", part2(graph))
}
