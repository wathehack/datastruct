package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertFirst(t *testing.T) {
	l := new(Dll)
	l.InsertFirst(1)
	assert.Equal(t, 1, l.Head.Data)
	l.InsertFirst(2)
	assert.Equal(t, 2, l.Head.Data)
	assert.Equal(t, 1, l.Head.Next.Data)
	assert.Equal(t, 2, l.Head.Next.Prev.Data)
}

func TestInsertAfter(t *testing.T) {
	l := new(Dll)
	err := l.InsertAfter(l.Head, 1)
	assert.NotNil(t, err)
	l.Head = &Node{1, nil, nil}
	err = l.InsertAfter(l.Head, 2)
	assert.Nil(t, err)
	assert.Equal(t, 1, l.Head.Data)
	assert.Equal(t, 2, l.Head.Next.Data)
	assert.Equal(t, 1, l.Head.Next.Prev.Data)
}

func TestInsertBefore(t *testing.T) {
	l := new(Dll)
	err := l.InsertBefore(l.Head, 1)
	assert.NotNil(t, err)
	l.Head = &Node{1, nil, nil}
	err = l.InsertBefore(l.Head, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, l.Head.Data)
	assert.Equal(t, 1, l.Head.Next.Data)
	assert.Equal(t, 2, l.Head.Next.Prev.Data)
}

func TestDelete(t *testing.T) {
	l := new(Dll)
	err := l.Delete(1)
	assert.NotNil(t, err)
	l.Head = &Node{1, nil, nil}
	l.InsertAfter(l.Head, 2)
	l.InsertAfter(l.Head, 3)
	err = l.Delete(100)
	assert.NotNil(t, err)
	err = l.Delete(1)
	assert.Nil(t, err)
	assert.Equal(t, 3, l.Head.Data)
	assert.Equal(t, 2, l.Head.Next.Data)
	assert.Equal(t, 3, l.Head.Next.Prev.Data)
}

func TestSearch(t *testing.T) {
	l := new(Dll)
	i, err := l.Search(1)
	assert.NotNil(t, err)
	assert.Equal(t, -1, i)
	l.Head = &Node{1, nil, nil}
	l.InsertAfter(l.Head, 2)
	l.InsertAfter(l.Head, 3)
	i, err = l.Search(100)
	assert.NotNil(t, err)
	assert.Equal(t, -1, i)
	i, err = l.Search(1)
	assert.Nil(t, err)
	assert.Equal(t, 0, i)
}
