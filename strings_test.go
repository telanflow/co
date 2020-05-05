package co

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplace(t *testing.T) {
	asserts := assert.New(t)

	asserts.Equal("origin", Replace(map[string]string{
		"123": "321",
	}, "origin"))

	asserts.Equal("321origin321", Replace(map[string]string{
		"123": "321",
	}, "123origin123"))
	asserts.Equal("321new321", Replace(map[string]string{
		"123":    "321",
		"origin": "new",
	}, "123origin123"))
}
