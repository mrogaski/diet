package diet

type Tree[T Integer] struct {
	Root *Node[T]
}

type Node[T Integer] struct {
	First T
	Last  T
}

// NewTree returns a pointer to a new discrete interval encoding tree.
func NewTree[T Integer]() *Tree[T] {
	return &Tree[T]{}
}

func (tree *Tree[T]) Insert(elem T) error {
	tree.Root = &Node[T]{First: elem, Last: elem}

	return nil
}
