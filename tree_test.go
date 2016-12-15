package tree

import (
	"testing"
	"fmt"
)

func TestAVLTreeGetRoot(t *testing.T) {
	var tree ITree = new(AVLTree)
	r := tree.GetRoot()
	n, ok := r.(AVLNode)
	if !ok {
		fmt.Println(n)
		t.Error()
	}
}

