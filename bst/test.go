package bst
//
//
//
//n := t.Root.Find(key)
//
//if n == nil {
//return false
//}
//
//// Is Root
//if n.Parent == nil {
//
//if n.Left
//
//if n.Left != nil {
//if t.Root = n.Left; n.Right != nil {
//t.Root.insert(*n.Right)
//}
//} else if n.Right != nil {
//if t.Root = n.Right; n.Left != nil {
//t.Root.insert(*n.Left)
//}
//} else {
//t.Root = nil
//}
//
//return true
//}
//
//var toMove *Node = nil
//
//if n.Parent.Right != nil && n.Parent.Right.key >= key {
//n.Parent.Right = nil
//} else {
//n.Parent.Left = nil
//}
//
//if n.Right != nil && n.Left != nil {
//n.Right.Parent = nil
//n.Left.Parent = nil
//n.Left.insert(*n.Right)
//toMove = n.Left
//} else if n.Right != nil {
//toMove = n.Right
//} else if n.Left != nil {
//toMove = n.Left
//}
//
//if toMove != nil {
//n.Parent.insert(*toMove)
//}
//
//n.Parent = nil
//n.Left = nil
//n.Right = nil
//
////if n.Left != nil {
////	n.Parent = n.Left
////	if n.Right != nil {
////		n.Parent.insert(*n.Right)
////	}
////} else if n.Right != nil {
////	n.Parent = n.Right
////	if n.Left != nil {
////		n.Parent.insert(*n.Left)
////	}
////} else {
////	//t.Root = nil
////}


/*

// Remove parenthood
	if n.Parent.Right != nil && n.Parent.Right.key >= n.key {
		n.Parent.Right = nil
	} else {
		n.Parent.Left = nil
	}

	n, merged := n.MergeRTL()

	if merged || n.Left != nil {
		n.Left.Parent = nil
		n.Parent.Insert(*n.Left)
	} else if n.Right != nil {
		n.Right.Parent = nil
		n.Parent.Insert(*n.Right)
	}


	type opt func(*Node) error

func (n *Node) Options(options ...opt) error {
	for _, opt := range options {
		err := opt(n)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParentNode(parentNode Node) opt {
	return func(n *Node) error {
		n.Parent = &parentNode
		return nil
	}
}

func RightNode(rightNode Node) opt {
	return func(n *Node) error {
		n.Right = &rightNode
		rightNode.Parent = n
		return nil
	}
}

func LeftNode(leftNode Node) opt {
	return func(n *Node) error {
		n.Left = &leftNode
		leftNode.Parent = n
		return nil
	}
}

*/