package hashtable

import (
	"errors"
	"fmt"
)

type Node struct {
	Key   string
	Value string
	Next  *Node
}

type H struct {
	Size  int
	Table []*Node
}

func HashTable(s int) *H {
	h := new(H)
	h.Size = s
	h.Table = make([]*Node, s, s)
	return h
}

func (h *H) HashValue(k string) (int, error) {
	if h.Size <= 0 {
		return -1, errors.New("Hash table size must be more than 0")
	}
	var s rune
	for _, v := range k {
		s += v
	}
	i := int(s) % h.Size
	return i, nil
}

func (h *H) AddNode(k string, v string) error {
	i, err := h.HashValue(k)
	if err != nil {
		return err
	}
	n := new(Node)
	n.Key, n.Value, n.Next = k, v, nil
	if h.Table[i] == nil {
		h.Table[i] = n
	} else {
		p := h.Table[i]
		for p.Next != nil {
			p = p.Next
		}
		p.Next = n
	}
	return nil
}

func (h *H) PrintTable() error {
	if h.Size <= 0 {
		return errors.New("Hash table size must be more than 0")
	}
	for i := 0; i < h.Size; i++ {
		fmt.Printf("%d", i)
		p := h.Table[i]
		if p == nil {
			fmt.Printf(" (nil)")
		}
		for p != nil {
			fmt.Printf(" (Key: %s, Value: %s)", p.Key, p.Value)
			p = p.Next
		}
		fmt.Printf("\n")
	}
	return nil
}

func (h *H) SearchNode(k string) (*Node, error) {
	if h.Size <= 0 {
		return nil, errors.New("Hash table size must be more than 0")
	}
	i, err := h.HashValue(k)
	if err != nil {
		return nil, err
	}
	p := h.Table[i]
	for p != nil {
		if p.Key == k {
			return p, nil
		}
		p = p.Next
	}
	return nil, errors.New("Item not found")
}

func (h *H) DeleteNode(k string) error {
	if h.Size <= 0 {
		return errors.New("Hash table size must be more than 0")
	}
	i, err := h.HashValue(k)
	if err != nil {
		return err
	}
	if h.Table[i] == nil {
		return errors.New("Item not found")
	} else if h.Table[i].Key == k && h.Table[i].Next == nil {
		h.Table[i] = nil
	} else if h.Table[i].Key == k && h.Table[i].Next != nil {
		h.Table[i] = h.Table[i].Next
	} else {
		p := h.Table[i]
		c := h.Table[i].Next
		for c != nil && c.Key != k {
			p = c
			c = c.Next
		}
		if c == nil {
			return errors.New("Item not found")
		} else {
			c = c.Next
			p.Next = c
		}
	}
	return nil
}
