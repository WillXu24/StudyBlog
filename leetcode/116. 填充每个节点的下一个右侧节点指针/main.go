package main

import "fmt"

func main() {
	tree := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	//bfs(tree)

	res := connect(tree)
	fmt.Println(res)
}

// Node 奇怪的树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func bfs(root *Node) {
	if root == nil {
		return
	}
	nodeQueue := []*Node{root}
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		fmt.Print(node.Val)
		nodeQueue = nodeQueue[1:]
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
		}
	}
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	nodeQueue := []*Node{root}
	for len(nodeQueue) != 0 {
		queueLength := len(nodeQueue)
		for i := 0; i < queueLength; i++ {
			if i == queueLength-1 {
				nodeQueue[i].Next = nil
			} else {
				nodeQueue[i].Next = nodeQueue[i+1]
			}
			if nodeQueue[i].Left != nil {
				nodeQueue = append(nodeQueue, nodeQueue[i].Left, nodeQueue[i].Right)
			}
		}
		nodeQueue = nodeQueue[queueLength:]
	}
	return root
}
