package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert.Equal(t, 1, Plus(1, 0), "they should be equal")
}

func TestPlus2(t *testing.T) {
	assert.Equal(t, 2, Plus(1, 1), "they should be equal")
}

func TestPlus3(t *testing.T) {
	assert.Equal(t, 3, Plus(1, 2), "they should be equal")
}

func TestPlus10(t *testing.T) {
	time.Sleep(1)
	assert.Equal(t, 10, Plus(5, 5), "they should be equal")
}

func TestPlus100(t *testing.T) {
	time.Sleep(1)
	assert.Equal(t, 100, Plus(50, 50), "they should be equal")
}
func TestPlus100Slow(t *testing.T) {
	time.Sleep(10)
	assert.Equal(t, 100, Plus(50, 50), "they should be equal")
}
