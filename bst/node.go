package bst

type Node struct {
	key    int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(key int) Node {
	return Node{key, nil, nil, nil}
}

func (n *Node) Insert(leaf Node) {
	if n == nil {
		return
	}
	current := n
	var parent *Node
	for current != nil {
		parent = current
		if leaf.key >= current.key {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	if leaf.key >= parent.key {
		parent.Right = &leaf
	} else {
		parent.Left = &leaf
	}
	leaf.Parent = parent
}

func (n *Node) Find(key int) *Node {
	for current := n; current != nil; {
		if current.key == key {
			return current
		}
		if key >= current.key {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return nil
}

func (n Node) RightFirst() []int {
	var sorted []int
	if n.Right != nil {
		sorted = append(sorted, n.Right.RightFirst()...)
	}
	sorted = append(sorted, n.key)
	if n.Left != nil {
		sorted = append(sorted, n.Left.RightFirst()...)
	}
	return sorted
}

func (n Node) RootFirst() []int {
	var sorted []int
	sorted = append(sorted, n.key)
	if n.Right != nil {
		sorted = append(sorted, n.Right.RootFirst()...)
	}
	if n.Left != nil {
		sorted = append(sorted, n.Left.RootFirst()...)
	}
	return sorted
}

func (n Node) LeftFirst() []int {
	var sorted []int
	if n.Left != nil {
		sorted = append(sorted, n.Left.LeftFirst()...)
	}
	sorted = append(sorted, n.key)
	if n.Right != nil {
		sorted = append(sorted, n.Right.LeftFirst()...)
	}
	return sorted
}

func (n Node) Closest() (node *Node, found bool) {
	if node, _ := n.ClosestRight(); node != nil {
		return node, true
	} else if node, _ := n.ClosestLeft(); node != nil {
		return node, true
	}
	return nil, false
}

func (n *Node) ClosestRight() (node *Node, found bool) {
	if n.Right == nil {
		return nil, false
	}
	current := n.Right
	for current != nil {
		if current.Left == nil {
			return current, true
		}
		current = current.Left
	}
	return current, true
}

func (n *Node) ClosestLeft() (node *Node, found bool) {
	if n.Left == nil {
		return nil, false
	}
	current := n.Left
	for current != nil {
		if current.Right == nil {
			return current, true
		}
		current = current.Right
	}
	return current, true
}

func (n *Node) MergeRTL() (node *Node, moved bool) {
	if n.Right == nil || n.Left == nil {
		return n, false
	}
	n.Right.Parent = nil
	right := n.Right
	n.Right = nil
	n.Left.Insert(*right)
	return n, true
}

func (n *Node) MergeLTR() (node *Node, moved bool) {
	if n.Right == nil || n.Left == nil {
		return n, false
	}
	n.Left.Parent = nil
	left := n.Left
	n.Left = nil
	n.Right.Insert(*left)
	return n, true
}

func (n *Node) Swap(toSwap *Node) bool {

	if toSwap == nil {
		return false
	}

	//copy := toSwap
	//toSwap.Remove()
	//
	//copy.Parent = n.Parent
	//copy.Left = n.Left
	//copy.Right = n.Right
	//
	//n = copy
	//
	//toSwap.Parent = n.Parent

	return true

}

func (n *Node) Remove() bool {



}
