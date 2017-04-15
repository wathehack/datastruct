package linkedlist

import (
	"errors"
)

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type Dll struct {
	Head *Node
}

func (l *Dll) InsertFirst(v interface{}) {
	n := new(Node)
	n.Data = v
	if l.Head == nil {
		n.Next = nil
		n.Prev = nil
		l.Head = n
	} else {
		n.Next = l.Head
		n.Prev = nil
		l.Head.Prev = n
		l.Head = n
	}
}

func (l *Dll) InsertAfter(n *Node, v interface{}) error {
	if n == nil {
		return errors.New("Insert after a nil node")
	}
	an := new(Node)
	an.Data = v
	an.Next = n.Next
	an.Prev = n
	if n.Next != nil {
		n.Next.Prev = an
	}
	n.Next = an
	return nil
}

func (l *Dll) InsertBefore(n *Node, v interface{}) error {
	if n == nil {
		return errors.New("Insert before a nil node")
	}
	bn := new(Node)
	bn.Data = v
	bn.Next = n
	bn.Prev = n.Prev
	if n.Prev != nil {
		n.Prev.Next = bn
	} else {
		l.Head = bn
	}
	n.Prev = bn
	return nil
}

func (l *Dll) Delete(v interface{}) error {
	n := l.Head
	if n == nil {
		return errors.New("Delete from an empty linked list")
	}
	for n != nil {
		if n.Data == v {
			if n.Next != nil {
				n.Next.Prev = n.Prev
			}
			if n.Prev != nil {
				n.Prev.Next = n.Next
			} else {
				l.Head = n.Next
			}
			n = nil
			return nil
		}
		n = n.Next
	}
	return errors.New("Item not found")
}

func (l *Dll) Search(v interface{}) (int, error) {
	n := l.Head
	if n == nil {
		return -1, errors.New("Search in an empty linked list")
	}
	c := 0
	for n != nil {
		if n.Data == v {
			return c, nil
		}
		n = n.Next
		c++
	}
	return -1, errors.New("Item not found")
}
