package stack

import (
	"errors"
)

type S []interface{}

func (s *S) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *S) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Pop from an empty stack")
	}
	v := (*s)[len(*s)-1]
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *S) Peek() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("Empty stack")
	}
	return (*s)[len(*s)-1], nil
}

func (s *S) Search(v interface{}) (int, error) {
	if len(*s) == 0 {
		return -1, errors.New("Search in an empty stack")
	}
	for i, e := range *s {
		if v == e {
			return i, nil
		}
	}
	return -1, errors.New("Item not found")
}

func (s *S) Size() int {
	return len(*s)
}

func (s *S) IsEmpty() bool {
	return len(*s) == 0
}
