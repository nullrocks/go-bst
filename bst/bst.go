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

	if t.Root.Key != key {
		n := t.Root.Find(key)
		if n == nil {
			return false
		}
		return n.Remove()
	}

	// Case 1: No children
	if t.Root.Left == nil && t.Root.Right == nil {
		t.Root = nil
		return true
	} else if t.Root.Left != nil && t.Root.Right != nil {
		// Case 2: Two children
		closest, _ := t.Root.ClosestRight() // Get the closest right node
		t.Root.Swap(closest)
		return closest.Remove()
	}

	// Case 3: One Child
	if t.Root.Right != nil {
		t.Root.Right.Parent = nil
		t.Root = t.Root.Right
	} else {
		t.Root.Left.Parent = nil
		t.Root = t.Root.Left
	}

	return true
}

func (t Tree) Find(key int) *Node {
	return t.Root.Find(key)
}

func (t Tree) InOrder() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.InOrder()
}

func (t Tree) PreOrder() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.PreOrder()
}

func (t Tree) PostOrder() []int {
	if t.Root == nil {
		return []int{}
	}
	return t.Root.PostOrder()
}
