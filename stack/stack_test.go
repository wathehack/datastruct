package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	var s S
	s.Push(5)
	s.Push(17)
	s.Push(32)
	assert.Equal(t, S{5, 17, 32}, s)
}

func TestPop(t *testing.T) {
	var s S
	v, err := s.Pop()
	assert.NotNil(t, err)
	s = S{5, 17, 32}
	v, err = s.Pop()
	assert.Equal(t, 32, v)
}

func TestPeek(t *testing.T) {
	var s S
	v, err := s.Peek()
	assert.NotNil(t, err)
	s = S{5, 17, 32}
	v, err = s.Peek()
	assert.Equal(t, 32, v)
}

func TestSearch(t *testing.T) {
	var s S
	v, err := s.Search(100)
	assert.NotNil(t, err)
	s = S{5, 17, 32}
	v, err = s.Search(100)
	assert.Equal(t, -1, v)
	v, err = s.Search(5)
	assert.Equal(t, 0, v)
}

func TestSize(t *testing.T) {
	var s S
	s = S{5, 17, 32}
	v := s.Size()
	assert.Equal(t, 3, v)
}

func TestIsEmpty(t *testing.T) {
	var s S
	v := s.IsEmpty()
	assert.Equal(t, true, v)
	s = S{5, 17, 32}
	v = s.IsEmpty()
	assert.Equal(t, false, v)
}
