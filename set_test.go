package goniq_test

import (
	"testing"

	"github.com/felbit/goniq"
	"github.com/stretchr/testify/assert"
)

func TestBasicOperations(t *testing.T) {
	set := goniq.NewSet[int]()
	assert.Equal(t, 0, set.Size())
	set.Add(1)          // set == [1]
	set.Append(2, 3, 1) // set == [1, 2, 3]
	assert.Equal(t, 3, set.Size())
	set.Add(3) // set == [1, 2, 3]
	assert.Equal(t, 3, set.Size())
	set.Remove(3) // set == [1, 2]
	assert.Equal(t, 2, set.Size())
	set.Remove(3) // set == [1, 2]
	assert.Equal(t, 2, set.Size())
	set.RemoveAll() // set == []
	assert.Equal(t, 0, set.Size())
}
