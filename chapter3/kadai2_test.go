package chapter3

import (
	"testing"

	"gotest.tools/assert"
)

func TestKadai2(t *testing.T) {
	k := NewKadai2(1, "hoge")

	assert.Equal(t, 1, k.id)
	assert.Equal(t, "hoge", k.name)
}
