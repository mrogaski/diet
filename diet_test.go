package diet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/go-diet"
)

func TestNewTree(t *testing.T) {
	t.Parallel()

	tree := diet.NewTree[int]()

	assert.IsType(t, &diet.Tree[int]{}, tree)
}

func TestTree_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		tree     *diet.Tree[int]
		elem     int
		expected bool
	}{
		{name: "empty", tree: &diet.Tree[int]{}, elem: 0, expected: false},
		{
			name:     "single element",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: 0, Last: 0}},
			elem:     0,
			expected: true,
		},
		{
			name:     "single node, internal",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 2}},
			elem:     0,
			expected: true,
		},
		{
			name:     "single node, left edge",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: 0, Last: 2}},
			elem:     0,
			expected: true,
		},
		{
			name:     "single node, right edge",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 0}},
			elem:     0,
			expected: true,
		},
		{
			name:     "single node, predecessor",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 2}},
			elem:     -3,
			expected: false,
		},
		{
			name:     "single node, successor",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 2}},
			elem:     3,
			expected: false,
		},
		{
			name:     "single node, less than",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 2}},
			elem:     -4,
			expected: false,
		},
		{
			name:     "single node, greater than",
			tree:     &diet.Tree[int]{Interval: &diet.Interval[int]{First: -2, Last: 2}},
			elem:     4,
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, tt.tree.Contains(tt.elem))
		})
	}
}

func TestTree_Insert(t *testing.T) {
	t.Skip("not implemented")
	t.Parallel()

	tests := []struct {
		name     string
		tree     *diet.Tree[int]
		input    []int
		expected *diet.Tree[int]
	}{
		{
			name:  "initial",
			tree:  &diet.Tree[int]{},
			input: []int{0},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: 0,
					Last:  0,
				},
			},
		},
		{
			name:  "duplicate",
			tree:  &diet.Tree[int]{},
			input: []int{0, 0},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: 0,
					Last:  0,
				},
			},
		},
		{
			name:  "predecessor",
			tree:  &diet.Tree[int]{},
			input: []int{0, -1},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: -1,
					Last:  0,
				},
			},
		},
		{
			name:  "successor",
			tree:  &diet.Tree[int]{},
			input: []int{0, 1},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: 0,
					Last:  1,
				},
			},
		},
		{
			name:  "left",
			tree:  &diet.Tree[int]{},
			input: []int{0, -5},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: 0,
					Last:  0,
				},
				Left: &diet.Tree[int]{
					Interval: &diet.Interval[int]{
						First: -5,
						Last:  -5,
					},
				},
			},
		},
		{
			name:  "right",
			tree:  &diet.Tree[int]{},
			input: []int{0, 5},
			expected: &diet.Tree[int]{
				Interval: &diet.Interval[int]{
					First: 0,
					Last:  0,
				},
				Right: &diet.Tree[int]{
					Interval: &diet.Interval[int]{
						First: 5,
						Last:  5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, elem := range tt.input {
				err := tt.tree.Insert(elem)
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, tt.tree)
		})
	}
}
