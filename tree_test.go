package tree

import (
	"testing"
	"fmt"
)

func TestAVLTreeGetRootIsAVLNode(t *testing.T) {
	var tree IBTree = new(AVLTree)
	r := tree.GetRoot()
	n, ok := r.(AVLNode)
	if !ok {
		fmt.Println(n)
		t.Error()
	}
}

func TestAVLNodeAddElement(t *testing.T) {
	var tree IBTree = new(AVLTree)
	tree.Add(AVLElement{Key:10})
	// n.Add()
	// n.Add(AVLElement{Key:15})
	// n.Add(AVLElement{Key:20})
	// n.GetElement().
}