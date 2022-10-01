// Package diet implements a discrete interval encoding tree in Go.
//
// The discrete interval encoding tree is a structure, described in [Diets for Fat Sets], for storing subsets
// of types having a total order and a predecessor and a successor function. The general idea is to represent
// a set by a binary search tree of integers in which maximal adjacent subsets are each represented by an interval.
//
// [Diets for Fat Sets]: https://web.engr.oregonstate.edu/~erwig/papers/abstracts.html#JFP98
package diet

// Tree represents the discrete interval encoding tree or a subtree.
type Tree[T Integer] struct {
	Interval *Interval[T] // subset represented by the node
	Left     *Tree[T]     // preceding subset
	Right    *Tree[T]     // succeeding subset
}

// Interval represents a subset consisting of a sequence of contiguous elements.
// Single-element subsets are represented by an interval where First and Last are equal.
type Interval[T Integer] struct {
	First T
	Last  T
}

// NewTree returns a pointer to a new discrete interval encoding tree.
func NewTree[T Integer]() *Tree[T] {
	return &Tree[T]{}
}

// Contains returns true if elem is a member of one of the subsets within the tree, false otherwise.
func (tree *Tree[T]) Contains(elem T) bool {
	t := tree // current subtree

	for {
		if t.Interval == nil {
			return false
		}

		switch {
		case elem >= t.Interval.First && elem <= t.Interval.Last:
			return true
		case elem < t.Interval.First:
			if t.Left == nil {
				t.Left = &Tree[T]{}
			}

			t = t.Left
		case elem > t.Interval.Last:
			if t.Right == nil {
				t.Right = &Tree[T]{}
			}

			t = t.Right
		default:
			return false
		}
	}
}

func (tree *Tree[T]) Insert(elem T) error {
	return nil
}
