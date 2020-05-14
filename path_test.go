package co

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathFillSlash(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/", PathFillSlash("/"))
	asserts.Equal("/", PathFillSlash(""))
	asserts.Equal("/123/", PathFillSlash("/123"))
}

func TestPathRemoveSlash(t *testing.T) {
	asserts := assert.New(t)
	asserts.Equal("/", PathRemoveSlash("/"))
	asserts.Equal("/123/1236", PathRemoveSlash("/123/1236"))
	asserts.Equal("/123/1236", PathRemoveSlash("/123/1236/"))
}
