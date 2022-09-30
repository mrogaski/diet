package diet

import "golang.org/x/exp/constraints"

type Tree[T constraints.Integer] struct {
	Root *Node[T]
}

func NewTree[T constraints.Integer]() *Tree[T] {
	return &Tree[T]{}
}

type Node[T constraints.Integer] struct{}
