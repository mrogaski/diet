package diet

// Interval represents a subset consisting of a sequence of contiguous elements which represent a discrete interval.
// Single-element subsets are represented by an interval where First and Last are equal.
type Interval[T Integer] struct {
	First T
	Last  T
}

// NewInterval returns a pointer to a new subset.
func NewInterval[T Integer](elem T) *Interval[T] {
	return &Interval[T]{First: elem, Last: elem}
}

func (i *Interval[T]) has(elem T) bool {
	return elem >= i.First && elem <= i.Last
}

func (i *Interval[T]) adjacent(other *Interval[T]) bool {
	return i.adjacentElement(other.First) || i.adjacentElement(other.Last)
}

func (i *Interval[T]) adjacentElement(elem T) bool {
	return i.leftAdjacentElement(elem) || i.rightAdjacentElement(elem)
}

func (i *Interval[T]) leftAdjacentElement(elem T) bool {
	return elem == i.First-1
}

func (i *Interval[T]) rightAdjacentElement(elem T) bool {
	return elem == i.Last+1
}

func (i *Interval[T]) merge(other *Interval[T]) {
	if i.First > other.First {
		i.First = other.First
	}

	if i.Last < other.Last {
		i.Last = other.Last
	}
}
