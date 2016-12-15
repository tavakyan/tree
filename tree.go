package tree

type IElement interface {

}

type INode interface {
	GetElement() IElement
}

type IBNode interface {
	GetElement() IElement
	GetLeft() IBNode
	GetRight() IBNode
}

type ITree interface {
	GetRoot() INode
}

type IBTree interface {
	GetRoot() IBNode
}

type AVLElement struct {
}

type AVLNode struct {
	element AVLElement
	left *AVLNode
	right *AVLNode
}

func (n AVLNode) GetElement() IElement {
	return n.element
}

func (n AVLNode) GetLeft() IBNode {
	return n.left
}

func (n AVLNode) GetRight() IBNode {
	return n.right
}

type AVLTree struct {
	root AVLNode
}

func (t AVLTree) GetRoot() IBNode {
	return t.root
}