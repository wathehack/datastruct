package graph

import (
	"errors"
	"fmt"

	"github.com/wathehack/datastruct/queue"
)

// The structure to represent an adjacency list node.
type Node struct {
	Dest int
	Next *Node
}

// The structure to represent an adjacency list.
type List struct {
	Head *Node
}

// The structure to represent a graph. Array stores all adjacency lists.
type G struct {
	Vertices int
	Array    []*List
}

// Create an adjacency list node.
func CreateNode(d int) (*Node, error) {
	if d < 0 {
		return nil, errors.New("Destination must be more than 0")
	}
	n := new(Node)
	n.Dest = d
	n.Next = nil
	return n, nil
}

// Create a graph of v vertices.
func CreateGraph(v int) (*G, error) {
	if v < 0 {
		return nil, errors.New("Number of vertices must be more than 0")
	}
	g := new(G)
	g.Vertices = v
	g.Array = make([]*List, v)
	for i := 0; i < v; i++ {
		g.Array[i] = new(List)
	}
	return g, nil
}

// Add an edge to an undirected graph.
func (g *G) AddEdge(s int, d int) error {
	if s < 0 || d < 0 {
		return errors.New("Source and destination must be more than 0")
	}
	// Add an edge from src to dest.
	n, err := CreateNode(d)
	if err != nil {
		return err
	}
	n.Next = g.Array[s].Head
	g.Array[s].Head = n
	// Add an edge from dest to src also since the graph is undirected.
	n, err = CreateNode(s)
	if err != nil {
		return err
	}
	n.Next = g.Array[d].Head
	g.Array[d].Head = n
	return nil
}

// Print the adjacenncy list representation of a graph.
func (g *G) PrintGraph() error {
	if g.Vertices < 0 {
		return errors.New("Number of vertices must be more than 0")
	}
	for i := 0; i < g.Vertices; i++ {
		fmt.Printf("%d\nhead", i)
		p := g.Array[i].Head
		for p != nil {
			fmt.Printf("->%d", p.Dest)
			p = p.Next
		}
		if p == nil {
			fmt.Printf("->nil")
		}
		fmt.Printf("\n")
	}
	return nil
}

// Breadth first search on a graph.
func (g *G) BreadthFirstSearch(s int) ([]int, error) {
	if g.Vertices < 0 {
		return nil, errors.New("Number of vertices must be more than 0")
	}
	// Mark all vertices as not visited.
	vd := make([]bool, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		vd[i] = false
	}
	q := new(queue.Q)
	res := make([]int, 0)
	// Mark the current node as visited and enqueue it.
	vd[s] = true
	q.Enqueue(s)
	for q.IsEmpty() == false {
		v, err := q.Dequeue()
		if err != nil {
			return nil, err
		}
		s = v.(int)
		res = append(res, s)
		p := g.Array[s].Head
		// Get all adjacent vertices of the dequeued vertex s. If an adjacent
		// vertex has not been visited, mark it visited and enqueue it.
		for p != nil {
			if vd[p.Dest] == false {
				vd[p.Dest] = true
				q.Enqueue(p.Dest)
			}
			p = p.Next
		}
	}
	return res, nil
}

// Depth first search on a graph.
func (g *G) DepthFirstSearch(s int) ([]int, error) {
	if g.Vertices < 0 {
		return nil, errors.New("Number of vertices must be more than 0")
	}
	// Mark all vertices as not visited.
	vd := make([]bool, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		vd[i] = false
	}
	res := make([]int, 0)
	var f func(s int)
	f = func(s int) {
		vd[s] = true
		res = append(res, s)
		p := g.Array[s].Head
		for p != nil {
			if vd[p.Dest] == false {
				f(p.Dest)
			}
			p = p.Next
		}
	}
	f(s)
	return res, nil
}
