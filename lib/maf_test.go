package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert.Equal(t, 2, Plus(1, 1), "they should be equal")
}
