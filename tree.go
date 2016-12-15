package tree

type IElement interface {

}

type INode interface {
	GetElement() IElement
}

type ITree interface {
	GetRoot() INode
}

type AVLElement struct {
}

type AVLNode struct {
	element AVLElement
}

func (n AVLNode) GetElement() IElement {
	return n.element
}

type AVLTree struct {
	root AVLNode
}

func (t AVLTree) GetRoot() INode {
	return t.root
}