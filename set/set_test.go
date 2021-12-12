package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeduplicate(t *testing.T) {
	elems := []interface{}{1, 1, 1, 2, 2, 3, 3, 4, 4}

	s := NewSet(elems...)

	assert.Equal(t, 4, s.Len())

	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
	assert.True(t, s.Has(3))
	assert.True(t, s.Has(4))
}
