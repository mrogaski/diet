package diet

type Tree[T Integer] struct {
	Root *Node[T]
}

type Node[T Integer] struct {
	First T
	Last  T
	Left  *Node[T]
	Right *Node[T]
}

// NewTree returns a pointer to a new discrete interval encoding tree.
func NewTree[T Integer]() *Tree[T] {
	return &Tree[T]{}
}

func (tree *Tree[T]) Insert(elem T) error {
	tree.Root = child(tree.Root, elem)
	node := tree.Root

	for {
		switch {
		case elem == node.First-1:
			node.First = elem

			return nil
		case elem < node.First:
			node.Left = child(node.Left, elem)
			node = node.Left
		case elem == node.Last+1:
			node.Last = elem

			return nil
		case elem > node.Last:
			node.Right = child(node.Right, elem)
			node = node.Right
		default:
			return nil
		}
	}
}

func child[T Integer](node *Node[T], elem T) *Node[T] {
	if node != nil {
		return node
	}

	return &Node[T]{First: elem, Last: elem}
}
