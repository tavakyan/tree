package tree

type ITree interface {
	GetRoot() INode
	Add(IElement)
}

type IBTree interface {
	GetRoot() IBNode
	Add(IElement)
}

type AVLTree struct {
	root AVLNode
}

func (t AVLTree) GetRoot() IBNode {
	return t.root
}

func (t AVLTree) Add(e IElement) {
	
}