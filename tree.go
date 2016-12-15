package tree

type INode interface {

}

type ITree interface {
	GetRoot() INode
}

type AVLNode struct {
}

type AVLTree struct {
	root AVLNode
}

func (t AVLTree) GetRoot() INode {
	return t.root
}