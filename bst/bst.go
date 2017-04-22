package bst

import (
	"errors"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type T struct {
	Root *Node
}

func (t *T) Insert(v int) {
	n := new(Node)
	n.Data = v
	if t.Root == nil {
		t.Root = n
		return
	}
	c := t.Root
	for {
		if v < c.Data {
			if c.Left == nil {
				c.Left = n
				return
			}
			c = c.Left
		} else {
			if c.Right == nil {
				c.Right = n
				return
			}
			c = c.Right
		}
	}
}

func (t *T) Delete(v int) error {
	if t.Root == nil {
		return errors.New("Delete from an empty binary search tree")
	}
	p := new(Node)
	c := t.Root
	for {
		if v == c.Data {
			break
		} else if v < c.Data {
			if c.Left == nil {
				return errors.New("Item not found")
			}
			p = c
			c = c.Left
		} else {
			if c.Right == nil {
				return errors.New("Item not found")
			}
			p = c
			c = c.Right
		}
	}
	if c.Right == nil {
		if c.Left != nil {
			c.Data = c.Left.Data
			c.Right = c.Left.Right
			c.Left = c.Left.Left
		} else {
			p.Right = nil
			c = nil
		}
		return nil
	}
	r := c
	c = c.Right
	for {
		if c.Left == nil {
			r.Data = c.Data
			c = nil
			return nil
		} else if c.Left != nil {
			c = c.Left
		}
	}
}

func (t *T) Search(v int) (*Node, error) {
	c := t.Root
	for {
		if v == c.Data {
			return c, nil
		} else if v < c.Data {
			if c.Left == nil {
				return nil, errors.New("Item not found")
			}
			c = c.Left
		} else {
			if c.Right == nil {
				return nil, errors.New("Item not found")
			}
			c = c.Right
		}
	}
}

func (t *T) SearchMin() (*Node, error) {
	if t.Root == nil {
		return nil, errors.New("Search in an empty binary search tree")
	}
	c := t.Root
	for {
		if c.Left == nil {
			return c, nil
		} else if c.Left != nil {
			c = c.Left
		}
	}
}

func (t *T) DeleteMin() error {
	if t.Root == nil {
		return errors.New("Search in an empty binary search tree")
	}
	p := new(Node)
	c := t.Root
	for {
		if c.Left == nil {
			p.Left = c.Right
			c = nil
			return nil
		} else if c.Left != nil {
			p = c
			c = c.Left
		}
	}
}
