package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	bst := &T{&Node{0, nil, nil}}
	bst.Insert(-10)
	bst.Insert(10)
	bst.Insert(-5)
	bst.Insert(5)
	assert.Equal(t, 0, bst.Root.Data)
	assert.Equal(t, -10, bst.Root.Left.Data)
	assert.Equal(t, 10, bst.Root.Right.Data)
	assert.Equal(t, -5, bst.Root.Left.Right.Data)
	assert.Equal(t, 5, bst.Root.Right.Left.Data)
}

func TestDelete(t *testing.T) {
	bst := &T{nil}
	err := bst.Delete(100)
	assert.NotNil(t, err)
	bst = &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}
	err = bst.Delete(100)
	assert.NotNil(t, err)
	err = bst.Delete(10)
	assert.Equal(t, &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{20, &Node{5, nil, nil}, nil}}}, bst)
	err = bst.Delete(5)
	assert.Equal(t, &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{20, nil, nil}}}, bst)
}

func TestSearch(t *testing.T) {
	bst := &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}
	n, err := bst.Search(100)
	assert.NotNil(t, err)
	assert.Nil(t, n)
	n, err = bst.Search(0)
	assert.Nil(t, err)
	assert.Equal(t, 0, n.Data)
	n, err = bst.Search(-5)
	assert.Nil(t, err)
	assert.Equal(t, -5, n.Data)
	n, err = bst.Search(5)
	assert.Nil(t, err)
	assert.Equal(t, 5, n.Data)
}

func TestSearchMin(t *testing.T) {
	bst := &T{nil}
	n, err := bst.SearchMin()
	assert.NotNil(t, err)
	bst = &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}
	n, err = bst.SearchMin()
	assert.Equal(t, -20, n.Data)
}

func TestDeleteMin(t *testing.T) {
	bst := &T{nil}
	err := bst.DeleteMin()
	assert.NotNil(t, err)
	bst = &T{&Node{0, &Node{-10, &Node{-20, nil, nil}, &Node{-5, nil, nil}}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}
	err = bst.DeleteMin()
	assert.Equal(t, &T{&Node{0, &Node{-10, nil, &Node{-5, nil, nil}}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}, bst)
	err = bst.DeleteMin()
	assert.Equal(t, &T{&Node{0, &Node{-5, nil, nil}, &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}}}, bst)
}
