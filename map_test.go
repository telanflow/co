package co

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeStringMap(t *testing.T) {
	ret := MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"A": "v1"}, false)
	assert.Equal(t, map[string]string{"A": "v0"}, ret)

	ret = MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"a": "v1"}, true)
	assert.Equal(t, map[string]string{"a": "v0"}, ret)
}

func TestKeyToLower(t *testing.T) {
	src := map[string]string{"A": "v0"}
	ret := KeyToLower(src)

	assert.Contains(t, ret, "a")
	assert.NotContains(t, ret, "A")
}

func TestGetByPath(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "val0",
		"key1": map[string]string{"sk0": "sv0"},
		"key2": []string{"sv1", "sv2"},
		"key3": map[string]interface{}{"sk1": "sv1"},
	}

	v, ok := GetByPath("key0", mp)
	assert.True(t, ok)
	assert.Equal(t, "val0", v)

	v, ok = GetByPath("key1.sk0", mp)
	assert.True(t, ok)
	assert.Equal(t, "sv0", v)

	v, ok = GetByPath("key3.sk1", mp)
	assert.True(t, ok)
	assert.Equal(t, "sv1", v)

	// not exists
	v, ok = GetByPath("not-exits", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
	v, ok = GetByPath("key2.not-exits", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
	v, ok = GetByPath("not-exits.subkey", mp)
	assert.False(t, ok)
	assert.Nil(t, v)

	// dont support array/slice
	v, ok = GetByPath("key2.1", mp)
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestKeys(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	ln := len(mp)
	ret := Keys(mp)
	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "key0")
	assert.Contains(t, ret, "key1")
	assert.Contains(t, ret, "key2")

	ret = Keys(&mp)
	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "key0")
	assert.Contains(t, ret, "key1")

	ret = Keys(struct {
		a string
	}{"v"})

	assert.Len(t, ret, 0)
}

func TestValues(t *testing.T) {
	mp := map[string]interface{}{
		"key0": "v0",
		"key1": "v1",
		"key2": 34,
	}

	ln := len(mp)
	ret := Values(mp)

	assert.Len(t, ret, ln)
	assert.Contains(t, ret, "v0")
	assert.Contains(t, ret, "v1")
	assert.Contains(t, ret, 34)

	ret = Values(struct {
		a string
	}{"v"})

	assert.Len(t, ret, 0)
}
