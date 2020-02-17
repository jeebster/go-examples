package mathematics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	multiSlice := []int{10, 5, 3, 6, 2, 4}
	singleSlice := []int{1}

	assert.Equal(t, 2, Min(multiSlice), "Min() should return the lowest slice value")
	assert.Equal(t, 1, Min(singleSlice), "Min() should return the lowest slice value")
}

func TestMax(t *testing.T) {
	multiSlice := []int{4, 1, 25, 57, 12, 10, 100, 0, 17}
	singleSlice := []int{10}

	assert.Equal(t, 100, Max(multiSlice), "Max() should return the largest slice value")
	assert.Equal(t, 10, Max(singleSlice), "Min() should return the largest slice value")
}
