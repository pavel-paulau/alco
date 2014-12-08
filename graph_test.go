package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
The file contains the adjacency list representation of a simple undirected
graph. There are 200 vertices labeled 1 to 200. The first column in the file
represents the vertex label, and the particular row (other entries except the
first column) tells all the vertices that the vertex is adjacent to.

So for example, the row looks like : "6 155 56 52 120 ......". This just means
that the vertex with label 6 is adjacent to (i.e., shares an edge with) the
vertices with labels 155,56,52,120,......,etc

The task is to code up and run the randomized contraction algorithm for the min
cut problem and use it on the above graph to compute the min cut (i.e., the
minimum-possible number of crossing edges).

HINT: Note that you'll have to figure out an implementation of edge
contractions. Initially, you might want to do this naively, creating a new graph
from the old every time there's an edge contraction. But you should also think
about more efficient implementations.

WARNING: Please make sure to run the algorithm many times with different random
seeds, and remember the smallest cut that you ever find.
*/
func TestRandomContraction(t *testing.T) {
	file, err := os.Open("data/kargerMinCut.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph := map[int64][]int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vertices := strings.Fields(scanner.Text())
		vertex, err := strconv.ParseInt(vertices[0], 10, 16)
		if err != nil {
			log.Fatal(err)
		}
		graph[vertex] = []int64{}
		for _, adjVertex := range vertices[1:] {
			adjVertexInt, err := strconv.ParseInt(adjVertex, 10, 16)
			if err != nil {
				log.Fatal(err)
			}
			graph[vertex] = append(graph[vertex], adjVertexInt)
		}
	}
	minCut := randomContraction(graph)
	assert.Equal(t, 17, minCut)
}
