package main

import (
	"fmt"
	"sync"
)

type Item interface {
}

type Stack struct {
	items []Item
	lock  sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		items: []Item{},
	}
}

func (s *Stack) Print() {
	fmt.Println(s.items)
}

func (s *Stack) Push(item Item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, item)
}

func (s *Stack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
