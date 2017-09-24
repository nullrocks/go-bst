package bst_test

import (
	"testing"
	"github.com/nullrocks/go-bst/bst"
)

func TestNode_Insert(t *testing.T) {
	n := bst.NewNode(2)
	insRight := bst.NewNode(3)
	insLeft := bst.NewNode(1)

	n.Insert(insRight)

	if n.Right == nil {
		t.Error("Right Node shouldn't be nil")
	}
	if n.Right.key != 3 {
		t.Error("Expected 3 got", n.Right.key)
	}

	n.Insert(insLeft)
	if n.Left == nil {
		t.Error("Left Node shouldn't be nil")
	}
	if n.Left.key != 1 {
		t.Error("Expected 1 got", n.Right.key)
	}

}
