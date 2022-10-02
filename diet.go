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
		case t.Interval.has(elem):
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

// Insert adds the element to a subset in the tree, merging subsets as necessary.
func (tree *Tree[T]) Insert(elem T) {
	t := tree

	for {
		if t.Interval == nil {
			t.Interval = &Interval[T]{First: elem, Last: elem}

			return
		}

		switch {
		case t.Interval.has(elem):
			return
		case t.Interval.leftAdjacentElement(elem):
			t.Interval.First = elem
			t.joinLeft()

			return
		case t.Interval.rightAdjacentElement(elem):
			t.Interval.Last = elem
			t.joinRight()

			return
		case elem < t.Interval.First:
			t.Left = ternary(t.Left == nil, &Tree[T]{}, t.Left)
			t = t.Left
		case elem > t.Interval.Last:
			t.Right = ternary(t.Right == nil, &Tree[T]{}, t.Right)
			t = t.Right
		default:
			return
		}
	}
}

func ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}

	return b
}

func (tree *Tree[T]) joinLeft() {
	if tree.Left == nil {
		return
	}

	parent := tree
	child := tree.Left

	for child.Right != nil {
		parent = child
		child = child.Right
	}

	if tree.Interval.adjacent(child.Interval) {
		tree.Interval.merge(child.Interval)

		if parent == tree {
			parent.Left = child.Left
		} else {
			parent.Right = nil
		}
	}
}

func (tree *Tree[T]) joinRight() {
	if tree.Right == nil {
		return
	}

	parent := tree
	child := tree.Right

	for child.Left != nil {
		parent = child
		child = child.Left
	}

	if tree.Interval.adjacent(child.Interval) {
		tree.Interval.merge(child.Interval)

		if parent == tree {
			parent.Right = child.Right
		} else {
			parent.Left = nil
		}
	}
}
