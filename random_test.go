package co

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func TestRandStr(t *testing.T) {
	asserts := assert.New(t)

	// 0 长度字符
	randStr := RandStr(0)
	asserts.Len(randStr, 0)

	// 16 长度字符
	randStr = RandStr(16)
	asserts.Len(randStr, 16)

	// 32 长度字符
	randStr = RandStr(32)
	asserts.Len(randStr, 32)

	//相同长度字符
	sameLenStr1 := RandStr(32)
	sameLenStr2 := RandStr(32)
	asserts.NotEqual(sameLenStr1, sameLenStr2)
}

func TestRandBytes(t *testing.T) {
	asserts := assert.New(t)

	// 0 长度字符
	randBytes := RandBytes(0)
	asserts.Len(randBytes, 0)

	// 16 长度字符
	randBytes = RandBytes(16)
	asserts.Len(randBytes, 16)

	// 32 长度字符
	randBytes = RandBytes(32)
	asserts.Len(randBytes, 32)

	//相同长度字符
	sameLenBytes1 := RandBytes(32)
	sameLenBytes2 := RandBytes(32)
	asserts.NotEqual(sameLenBytes1, sameLenBytes2)
}

func TestRandStrToUpper(t *testing.T) {
	str := RandStrToUpper("abcdefghijklmnopqrstuvwxyz", 5)

	upper := 0
	for _, r := range []rune(str) {
		if unicode.IsUpper(r) {
			upper++
		}
	}

	asserts := assert.New(t)
	asserts.Equal(upper, 5)
}

func TestRandInt(t *testing.T) {
	n := RandInt(-200, 100)
	n1 := RandInt(-200, 100)
	asserts := assert.New(t)
	asserts.NotEqual(n, n1)
}
