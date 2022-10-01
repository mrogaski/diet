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

func TestTree_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		tree     *diet.Tree[int]
		input    []int
		expected *diet.Tree[int]
	}{
		{
			name:     "initial",
			tree:     &diet.Tree[int]{},
			input:    []int{0},
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: 0, Last: 0}},
		},
		{
			name:     "duplicate",
			tree:     &diet.Tree[int]{},
			input:    []int{0, 0},
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: 0, Last: 0}},
		},
		{
			name:     "predecessor",
			tree:     &diet.Tree[int]{},
			input:    []int{0, -1},
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: -1, Last: 0}},
		},
		{
			name:     "successor",
			tree:     &diet.Tree[int]{},
			input:    []int{0, 1},
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: 0, Last: 1}},
		},
		{
			name:  "left",
			tree:  &diet.Tree[int]{},
			input: []int{0, -5},
			expected: &diet.Tree[int]{
				Root: &diet.Node[int]{
					First: 0,
					Last:  0,
					Left: &diet.Node[int]{
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
				Root: &diet.Node[int]{
					First: 0,
					Last:  0,
					Right: &diet.Node[int]{
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
