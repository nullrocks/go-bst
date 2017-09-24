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

func (n Node) PreOrder() []int {
	var sorted []int
	sorted = append(sorted, n.key)
	if n.Right != nil {
		sorted = append(sorted, n.Right.PreOrder()...)
	}
	if n.Left != nil {
		sorted = append(sorted, n.Left.PreOrder()...)
	}
	return sorted
}

func (n Node) InOrder() []int {
	var sorted []int
	if n.Left != nil {
		sorted = append(sorted, n.Left.InOrder()...)
	}
	sorted = append(sorted, n.key)
	if n.Right != nil {
		sorted = append(sorted, n.Right.InOrder()...)
	}
	return sorted
}

func (n Node) PostOrder() []int {
	var sorted []int
	if n.Right != nil {
		sorted = append(sorted, n.Right.PostOrder()...)
	}
	sorted = append(sorted, n.key)
	if n.Left != nil {
		sorted = append(sorted, n.Left.PostOrder()...)
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

func (n Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n Node) IsRoot() bool {
	return n.Parent == nil
}

func (n Node) IsAncestorOf(child Node) (isParent bool, depth int) {
	return child.IsChildOf(n)
}

func (n Node) IsChildOf(parent Node) (isChild bool, depth int) {
	for np, d := n.Parent, 1; np == nil; {
		if np == &parent {
			return true, d
		}
		np = np.Parent
	}
	return false, 0
}

func (n *Node) Swap(toSwap *Node) bool {

	if toSwap == nil {
		return false
	}

	tmp := toSwap.key
	toSwap.key = n.key
	n.key = tmp

	return true

}

func (n *Node) SwapLeaf(leaf *Node) bool {
	if n == nil || leaf == nil || !leaf.IsLeaf() {
		return false
	}
	n.key = leaf.key
	return leaf.Remove()
}

func (n *Node) Remove() bool {
	isRightNode := false
	if n.Parent == nil {
		return false
	} else if n.Parent.Right != nil && n.Parent.Right.key == n.key {
		isRightNode = true
	}

	// Case 1: No children
	if n.Left == nil && n.Right == nil {
		if isRightNode {
			n.Parent.Right = nil
		} else {
			n.Parent.Left = nil
		}
		return true
	} else if n.Left != nil && n.Right != nil {
		// Case 2: Two children
		closest, _ := n.ClosestRight() // Get the closest right node
		n.Swap(closest)
		return closest.Remove()
	}

	// Case 3: One Child
	if n.Right != nil {
		n.Right.Parent = n.Parent
		if isRightNode {
			n.Parent.Right = n.Right
		} else {
			n.Parent.Left = n.Right
		}

	} else {
		n.Left.Parent = n.Parent
		if isRightNode {
			n.Parent.Right = n.Left
		} else {
			n.Parent.Left = n.Left
		}
	}

	return true

}
