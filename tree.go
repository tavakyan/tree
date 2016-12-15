package tree

type ITree interface {
	GetRoot() INode
}

type IBTree interface {
	GetRoot() IBNode
}

type AVLTree struct {
	root AVLNode
}

func (t AVLTree) GetRoot() IBNode {
	return t.root
}
