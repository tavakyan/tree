package tree

import (
	"testing"
	"fmt"
)

func TestAVLTreeGetRootIsAVLNode(t *testing.T) {
	var tree ITree = new(AVLTree)
	r := tree.GetRoot()
	n, ok := r.(AVLNode)
	if !ok {
		fmt.Println(n)
		t.Error()
	}
}

// func TestAVLNodeGetElementIsAVLElement(t *testing.T) {
// 	var node INode = new(AVLNode)
// 	e := node.GetElement()
// 	e1, ok := e.(AVLElement)
// 	if !ok {
// 		fmt.Println(e1)
// 		t.Error()
// 	}
// }