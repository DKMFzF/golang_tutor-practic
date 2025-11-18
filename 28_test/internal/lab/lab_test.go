package lab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, Sum(2, 2), 4)
	assert.Equal(t, Sum(5, 2), 7)
}
