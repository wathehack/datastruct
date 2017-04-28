package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	n, err := CreateNode(-1)
	assert.NotNil(t, err)
	n, err = CreateNode(1)
	assert.Nil(t, err)
	assert.Equal(t, &Node{1, nil}, n)
}

func TestCreateGraph(t *testing.T) {
	g, err := CreateGraph(-1)
	assert.NotNil(t, err)
	g, err = CreateGraph(1)
	assert.Nil(t, err)
	assert.Equal(t, &G{1, []*List{&List{nil}}}, g)
}

func TestAddEdge(t *testing.T) {
	g, err := CreateGraph(5)
	assert.Nil(t, err)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	assert.Equal(t, 2, g.Array[0].Head.Dest)
	assert.Equal(t, 1, g.Array[0].Head.Next.Dest)
	assert.Equal(t, 2, g.Array[1].Head.Dest)
	assert.Equal(t, 0, g.Array[1].Head.Next.Dest)
	assert.Equal(t, 1, g.Array[2].Head.Dest)
	assert.Equal(t, 0, g.Array[2].Head.Next.Dest)
}

func TestPrintGraph(t *testing.T) {
	g, err := CreateGraph(5)
	assert.Nil(t, err)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	err = g.PrintGraph()
	assert.Nil(t, err)
}

func TestBreadthFirstSearch(t *testing.T) {
	g, err := CreateGraph(10)
	assert.Nil(t, err)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 7)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(7, 9)
	g.AddEdge(8, 9)
	res, err := g.BreadthFirstSearch(0)
	assert.Nil(t, err)
	assert.Equal(t, []int{0, 2, 1, 3, 5, 4, 8, 7, 9, 6}, res)
}

func TestDepthFirstSearch(t *testing.T) {
	g, err := CreateGraph(10)
	assert.Nil(t, err)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 7)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(7, 9)
	g.AddEdge(8, 9)
	res, err := g.DepthFirstSearch(0)
	assert.Nil(t, err)
	assert.Equal(t, []int{0, 2, 3, 4, 7, 9, 8, 5, 1, 6}, res)
}
