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

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt(0, 100)
	}
}

func TestRandStrToUpperWithIndex(t *testing.T) {
	str := RandStrToUpperWithIndex("abcdefghijklmnopqrstuvwxyz", 1, 5, 10)
	asserts := assert.New(t)
	asserts.Equal(int32(str[1]), 'B')
	asserts.Equal(int32(str[5]), 'F')
	asserts.Equal(int32(str[10]), 'K')
}

func TestRandStrToLowerWithIndex(t *testing.T) {
	str := RandStrToLowerWithIndex("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, 5, 10)
	asserts := assert.New(t)
	asserts.Equal(int32(str[1]), 'b')
	asserts.Equal(int32(str[5]), 'f')
	asserts.Equal(int32(str[10]), 'k')
}
