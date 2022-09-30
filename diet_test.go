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
