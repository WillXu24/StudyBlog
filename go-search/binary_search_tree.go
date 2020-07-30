package main

type Node struct {
	value int
	left *Node
	right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {
	node:=&Node{
		value: value,
		left:  nil,
		right: nil,
	}

	if bst.Root==nil{
		bst.Root=node
		return
	}

	insertNode(bst.Root,node)
}

func insertNode(root,node *Node ) {
	if node.value>root.value{
		if root.left==nil{
			root.left=node
			return
		}
		insertNode(root.left,node)
	}
	if node.value>root.value{
		if root.right==nil{
			root.right=node
			return
		}
		insertNode(root.right,node)
	}
}

func (bst *BST) Remove (value int)  {
	
}