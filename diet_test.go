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
		elem     int
		expected *diet.Tree[int]
	}{
		{
			name:     "zero",
			tree:     &diet.Tree[int]{},
			elem:     0,
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: 0, Last: 0}},
		},
		{
			name:     "positive",
			tree:     &diet.Tree[int]{},
			elem:     5,
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: 5, Last: 5}},
		},
		{
			name:     "negative",
			tree:     &diet.Tree[int]{},
			elem:     -5,
			expected: &diet.Tree[int]{Root: &diet.Node[int]{First: -5, Last: -5}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.tree.Insert(tt.elem)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, tt.tree)
		})
	}
}
