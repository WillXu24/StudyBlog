package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) *Node {
	node := &Node{
		value: value,
		left:  nil,
		right: nil,
	}

	if bst.Root == nil {
		bst.Root = node
		return bst.Root
	}

	return insertNode(bst.Root, node)
}

func insertNode(root, node *Node) *Node {
	if node.value < root.value {
		if root.left == nil {
			root.left = node
			return root.left
		}
		return insertNode(root.left, node)
	}

	if node.value > root.value {
		if root.right == nil {
			root.right = node
			return root.right
		}
		return insertNode(root.right, node)
	}
	return root
}

func (bst *BST) Remove(value int) bool {
	node := bst.Search(value)
	if node == nil {
		return false
	}

	if node.right == nil && node.left == nil {
		node = nil
		return true
	}

	if node.left == nil {
		node = node.right
		return true
	}

	if node.right == nil {
		node = node.left
		return true
	}

	// todo: remove
}

func (bst *BST) Search(value int) *Node {
	if bst.Root == nil {
		return nil
	}

	node := bst.Root
	for node != nil {
		if value == node.value {
			return node
		}

		if value < node.value {
			node = node.left
			continue
		}
		if value > node.value {
			node = node.right
			continue
		}
	}

	return nil
}
