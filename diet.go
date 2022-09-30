package diet

type Tree[T Integer] struct {
	Root *Node[T]
}

func NewTree[T Integer]() *Tree[T] {
	return &Tree[T]{}
}

type Node[T Integer] struct{}
