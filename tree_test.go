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
