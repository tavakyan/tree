package tree

type INode interface {
	GetElement() IElement
}

type IBNode interface {
	GetElement() IElement
	GetLeft() IBNode
	GetRight() IBNode
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
