package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashValue(t *testing.T) {
	h := HashTable(0)
	v, err := h.HashValue("Paul")
	assert.NotNil(t, err)
	assert.Equal(t, -1, v)
	h = HashTable(10)
	v, err = h.HashValue("Kim")
	assert.Nil(t, err)
	assert.Equal(t, 9, v)
	v, err = h.HashValue("Emma")
	assert.Nil(t, err)
	assert.Equal(t, 4, v)
}

func TestAddNode(t *testing.T) {
	h := HashTable(10)
	h.AddNode("Paul", "86")
	h.AddNode("Kim", "72")
	h.AddNode("Emma", "98")
	h.AddNode("Annie", "18")
	h.AddNode("Sarah", "77")
	h.AddNode("Pepper", "52")
	h.AddNode("Mike", "90")
	h.AddNode("Steve", "24")
	h.AddNode("Bill", "55")
	h.AddNode("Marie", "68")
	h.AddNode("Susan", "85")
	h.AddNode("Joe", "93")
	assert.Equal(t, "86", h.Table[2].Value)
	assert.Equal(t, "72", h.Table[9].Value)
	assert.Equal(t, "98", h.Table[4].Value)
	assert.Equal(t, "72", h.Table[9].Value)
	assert.Equal(t, "24", h.Table[9].Next.Value)
}

func TestPrintTable(t *testing.T) {
	h := HashTable(0)
	err := h.PrintTable()
	assert.NotNil(t, err)
	h = HashTable(10)
	h.AddNode("Paul", "86")
	h.AddNode("Kim", "72")
	h.AddNode("Emma", "98")
	h.AddNode("Annie", "18")
	h.AddNode("Sarah", "77")
	h.AddNode("Pepper", "52")
	h.AddNode("Mike", "90")
	h.AddNode("Steve", "24")
	h.AddNode("Bill", "55")
	h.AddNode("Marie", "68")
	h.AddNode("Susan", "85")
	h.AddNode("Joe", "93")
	err = h.PrintTable()
	assert.Nil(t, err)
}

func TestSearchNode(t *testing.T) {
	h := HashTable(0)
	n, err := h.SearchNode("Paul")
	assert.NotNil(t, err)
	h = HashTable(10)
	n, err = h.SearchNode("Paul")
	assert.Nil(t, n)
	assert.NotNil(t, err)
	h.AddNode("Paul", "86")
	h.AddNode("Kim", "72")
	h.AddNode("Emma", "98")
	h.AddNode("Annie", "18")
	h.AddNode("Sarah", "77")
	h.AddNode("Pepper", "52")
	h.AddNode("Mike", "90")
	h.AddNode("Steve", "24")
	h.AddNode("Bill", "55")
	h.AddNode("Marie", "68")
	h.AddNode("Susan", "85")
	h.AddNode("Joe", "93")
	n, err = h.SearchNode("Paul")
	assert.Equal(t, "86", n.Value)
	assert.Nil(t, err)
	n, err = h.SearchNode("Steve")
	assert.Equal(t, "24", n.Value)
	assert.Nil(t, err)
}

func TestDeleteNode(t *testing.T) {
	h := HashTable(0)
	err := h.DeleteNode("Paul")
	assert.NotNil(t, err)
	h = HashTable(10)
	err = h.DeleteNode("Paul")
	assert.NotNil(t, err)
	h.AddNode("Paul", "86")
	h.AddNode("Kim", "72")
	h.AddNode("Emma", "98")
	h.AddNode("Annie", "18")
	h.AddNode("Sarah", "77")
	h.AddNode("Pepper", "52")
	h.AddNode("Mike", "90")
	h.AddNode("Steve", "24")
	h.AddNode("Bill", "55")
	h.AddNode("Marie", "68")
	h.AddNode("Susan", "85")
	h.AddNode("Joe", "93")
	err = h.DeleteNode("Annie")
	assert.Nil(t, err)
	err = h.DeleteNode("Emma")
	assert.Nil(t, err)
	err = h.DeleteNode("Mike")
	assert.Nil(t, err)
	err = h.PrintTable()
	assert.Nil(t, err)
}