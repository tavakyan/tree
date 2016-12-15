package tree

import (
	"testing"
	"fmt"
)

func TestAVLNodeGetElementIsAVLElement(t *testing.T) {
	var n INode = new(AVLNode)
	e := n.GetElement()
	e1, ok := e.(AVLElement)
	if !ok {
		fmt.Println(e1)
		t.Error()
	}
}

func TestAVLNodeGetLeftIsAVLNode(t *testing.T) {
	var n IBNode = new(AVLNode)
	l := n.GetLeft()
	l1, ok := l.(*AVLNode)
	if !ok {
		fmt.Println(l1)
		t.Error()
	}
}

func TestAVLNodeGetRightIsAVLNode(t *testing.T) {
	var n IBNode = new(AVLNode)
	r := n.GetRight()
	r1, ok := r.(*AVLNode)
	if !ok {
		fmt.Println(r1)
		t.Error()
	}
}

// func TestAVLNodeAddElement(t *testing.T) {
// 	var n IBNode = new(AVLNode)
// 	n.Add(AVLElement{Key:10})
// 	n.Add(AVLElement{Key:15})
// 	n.Add(AVLElement{Key:20})
// 	n.GetElement().
// }