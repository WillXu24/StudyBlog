package main

import (
	"fmt"
	"sync"
)

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				//Left:  nil,
				//Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
					//Left:  nil,
					//Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					//Left:  nil,
					//Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
				//Left:  nil,
				//Right: nil,
			},
			//Right: nil,
		},
	}

	fmt.Println(">>> Recurse <<<")
	PreorderRecurse(tree)
	fmt.Println()
	InorderRecurse(tree)
	fmt.Println()
	PostorderRecurse(tree)
	fmt.Println()

	fmt.Println(">>> DFS <<<")
	PreorderDFS(tree)
	fmt.Println()
	InorderDFS(tree)
	fmt.Println()
	PostorderDFS(tree)
	fmt.Println()

	fmt.Println(">>> BFS <<<")
	LevelOrderBFS(tree)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderRecurse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreorderRecurse(root.Left)
	PreorderRecurse(root.Right)
}

func InorderRecurse(root *TreeNode) {
	if root == nil {
		return
	}
	InorderRecurse(root.Left)
	fmt.Printf("%d ", root.Val)
	InorderRecurse(root.Right)
}

func PostorderRecurse(root *TreeNode) {
	if root == nil {
		return
	}
	PostorderRecurse(root.Left)
	PostorderRecurse(root.Right)
	fmt.Printf("%d ", root.Val)
}

func PreorderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	stack := NewStack()

	stack.Push(root)

	for !stack.Empty() {
		node := stack.Pop().(*TreeNode)
		fmt.Printf("%d ", node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

func InorderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	stack := NewStack()
	for {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		if stack.Empty() {
			break
		}

		root = stack.Pop().(*TreeNode)
		fmt.Printf("%d ", root.Val)
		root = root.Right
	}
}

func PostorderDFS(root *TreeNode) {
	if root == nil {
		return
	}

	stackA := NewStack()
	stackB := NewStack()

	stackA.Push(root)

	for !stackA.Empty() {
		node := stackA.Pop().(*TreeNode)
		stackB.Push(node)
		//fmt.Printf("%d ", node.Val)

		if node.Left != nil {
			stackA.Push(node.Left)
		}
		if node.Right != nil {
			stackA.Push(node.Right)
		}
	}

	for !stackB.Empty() {
		node := stackB.Pop().(*TreeNode)
		fmt.Printf("%d ", node.Val)
	}
}

func LevelOrderBFS(root *TreeNode) {
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		fmt.Printf("%d ", node.Val)

		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
		}
	}
}

/*******Stack******************************************/
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

func (s *Stack) Empty() bool {
	return len(s.items) == 0
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
