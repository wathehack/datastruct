package queue

import (
	"errors"
)

type Q []interface{}

func (q *Q) Enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Q) Dequeue() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("Dequeue from an empty queue")
	}
	v := (*q)[0]
	(*q)[0] = nil
	*q = (*q)[1:]
	return v, nil
}

func (q *Q) Peek() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("Empty queue")
	}
	return (*q)[0], nil
}

func (q *Q) Search(v interface{}) (int, error) {
	if len(*q) == 0 {
		return -1, errors.New("Search in an empty queue")
	}
	for i, e := range *q {
		if v == e {
			return i, nil
		}
	}
	return -1, errors.New("Item not found")
}

func (q *Q) Size() int {
	return len(*q)
}

func (q *Q) IsEmpty() bool {
	return len(*q) == 0
}
