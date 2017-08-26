package bst

type Tree struct {
	Root *Node
}

func NewTree(keys ...int) Tree {
	t := Tree{nil}
	for _, key := range keys {
		t.Insert(key)
	}
	return t
}

func (t *Tree) Insert(key int) {
	n := NewNode(key)
	if t.Root == nil {
		t.Root = &n
		return
	}
	t.Root.Insert(n)
}

func (t *Tree) Remove(key int) bool {

	if t == nil || t.Root == nil {
		return false
	}

	n := t.Find(key)

	if n == nil {
		return false
	}

	var container string

	if n.key == t.Root.key {
		container = "root"
	} else if n.Parent.Right != nil && n.Parent.Right.key == n.key {
		container = "right"
	} else {
		container = "left"
	}

	// Case 1: No children
	if n.Left == nil && n.Right == nil {
		switch container {
		case "root":
			t.Root = nil
		case "right":
			n.Parent.Right = nil
		case "left":
			n.Parent.Left = nil
		}
		return true
	} else if n.Left != nil && n.Right != nil {
		// Case 2: Two children
		closest, _ := n.ClosestRight() // Get the closest right node

		//closest.Remove()               // Remove the closest node to prevent duplicates
		//closest.Parent = n.Parent      // Rewire relations

		switch container {
		case "root":
			t.Root = nil
		case "right":
			n.Parent.Right = nil
		case "left":
			n.Parent.Left = nil
		}

		if rightNode {
			n.Parent.Right = closest
		} else {
			n.Parent.Left = closest
		}

		return true
	}

	// Case 3: One Child
	var onlyChild *Node

	if n.Right != nil {
		onlyChild = n.Right
	} else {
		onlyChild = n.Left
	}

	onlyChild.Parent = n.Parent
	if rightNode {
		n.Parent.Right = onlyChild
	} else {
		n.Parent.Left = onlyChild
	}

	return true
}

func (t Tree) Find(key int) *Node {
	return t.Root.Find(key)
}

func (t Tree) LeftFirst() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.LeftFirst()
}

func (t Tree) RootFirst() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.RootFirst()
}

func (t Tree) RightFirst() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.RightFirst()
}

//
//func (t Tree) Traverse() func(t Tree) [] int {
//
//}
