package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	var q Q
	q.Enqueue(5)
	q.Enqueue(17)
	q.Enqueue(32)
	assert.Equal(t, Q{5, 17, 32}, q)
}

func TestDequeue(t *testing.T) {
	var q Q
	v, err := q.Dequeue()
	assert.NotNil(t, err)
	q = Q{5, 17, 32}
	v, err = q.Dequeue()
	assert.Equal(t, 5, v)
}

func TestPeek(t *testing.T) {
	var s Q
	v, err := s.Peek()
	assert.NotNil(t, err)
	s = Q{5, 17, 32}
	v, err = s.Peek()
	assert.Equal(t, 5, v)
}

func TestSearch(t *testing.T) {
	var q Q
	v, err := q.Search(100)
	assert.NotNil(t, err)
	q = Q{5, 17, 32}
	v, err = q.Search(100)
	assert.Equal(t, -1, v)
	v, err = q.Search(5)
	assert.Equal(t, 0, v)
}

func TestSize(t *testing.T) {
	var s Q
	s = Q{5, 17, 32}
	v := s.Size()
	assert.Equal(t, 3, v)
}

func TestIsEmpty(t *testing.T) {
	var s Q
	v := s.IsEmpty()
	assert.Equal(t, true, v)
	s = Q{5, 17, 32}
	v = s.IsEmpty()
	assert.Equal(t, false, v)
}
