package main

import (
	"fmt"
)

type Graph struct {
	list map[int]map[int]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		list: make(map[int]map[int]struct{}),
	}
}

func (g *Graph) AddEdgeNotDirected(v1, v2 int) {
	if g.list[v1] == nil {
		g.list[v1] = make(map[int]struct{})
	}

	g.list[v1][v2] = struct{}{}

	if g.list[v2] == nil {
		g.list[v2] = make(map[int]struct{})
	}

	g.list[v2][v1] = struct{}{}
}

func (g *Graph) ShowGraph() {
	for vertex, neighbors := range g.list {
		fmt.Printf("Vertex %d -> %v\n", vertex, neighbors)
	}
}

func (g *Graph) EdgeExists(v1, v2 int) bool {
	neighbors, v1Exists := g.list[v1]
	if !v1Exists {
		return false
	}

	_, edgeExists := neighbors[v2]
	return edgeExists
}

func IsClique(g *Graph, vertexList []int) bool {
	for i := 0; i < len(vertexList); i++ {
		for j := i + 1; j < len(vertexList); j++ {
			if !g.EdgeExists(vertexList[i], vertexList[j]) {
				return false
			}
		}
	}

	return true
}

// Bounced Search Tree to define the vertex list
func Combinations(g *Graph, neighbors []int, kMinusOne, init, v int, currentCombination []int, results *[][]int) {
	if len(*results) > 0 {
		return
	}

	if len(currentCombination) == kMinusOne {
		if IsClique(g, currentCombination) {
			newCombination := make([]int, len(currentCombination))
			copy(newCombination, currentCombination)
			newCombination = append(newCombination, v)

			*results = append(*results, newCombination)
		}

		return
	}

	for i := init; i < len(neighbors); i++ {
		currentCombination = append(currentCombination, neighbors[i])
		Combinations(g, neighbors, kMinusOne, i+1, v, currentCombination, results)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func FPTClique(g *Graph, k int) string {
	maxDegree := 0
	for _, neighbors := range g.list {
		if len(neighbors) > maxDegree {
			maxDegree = len(neighbors)
		}
	}

	// Using the maxDegree parameter to prune the search space
	if k > maxDegree+1 {
		return "NO"
	}

	for vertex, neighborsMap := range g.list {
		// if the degree of the current vertex is less than k-1,
		// it cannot be part of a clique of size k
		if len(neighborsMap) >= k-1 {
			var neighborsSlice []int
			for neighborId := range neighborsMap {
				neighborsSlice = append(neighborsSlice, neighborId)
			}

			var results [][]int
			var currentCombination []int

			// Creating combinations of size k-1 from the neighbors of the current vertex
			Combinations(g, neighborsSlice, k-1, 0, vertex, currentCombination, &results)

			if len(results) > 0 {
				return "YES"
			}
		}
	}

	return "NO"
}

func main() {
	g := NewGraph()
	g.AddEdgeNotDirected(1, 2)
	g.AddEdgeNotDirected(1, 3)
	g.AddEdgeNotDirected(1, 4)
	g.AddEdgeNotDirected(2, 3)
	g.AddEdgeNotDirected(2, 4)
	g.AddEdgeNotDirected(3, 4)

	g.AddEdgeNotDirected(4, 5)
	g.AddEdgeNotDirected(5, 6)

	fmt.Println("Procurando Clique k=4:", FPTClique(g, 4)) // Expected: YES
	fmt.Println("Procurando Clique k=5:", FPTClique(g, 5)) // Expected: NO
}
