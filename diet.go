package diet

type Tree[T Integer] struct {
	Interval *Interval[T]
	Left     *Tree[T]
	Right    *Tree[T]
}

type Interval[T Integer] struct {
	First T
	Last  T
}

// NewTree returns a pointer to a new discrete interval encoding tree.
func NewTree[T Integer]() *Tree[T] {
	return &Tree[T]{}
}

func (tree *Tree[T]) Contains(elem T) bool {
	if tree.Interval == nil {
		return false
	}

	return elem >= tree.Interval.First && elem <= tree.Interval.Last
}

func (tree *Tree[T]) Insert(elem T) error {
	return nil
}
