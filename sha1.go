package co

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// 计算字符串的sha1散列值
func Sha1Str(s string) string {
	data := []byte(s)
	m := sha1.New()
	_, _ = m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

// 计算字符串的 shaX 散列值
// x为1/256/512
func ShaXStr(str string, x uint16) string {
	return string(shaXStr([]byte(str), x))
}

// 计算字符串的 shaX 散列值
// x为1/256/512
func shaXStr(str []byte, x uint16) []byte {
	var h hash.Hash
	switch x {
	case 1:
		h = sha1.New()
	case 256:
		h = sha256.New()
	case 512:
		h = sha512.New()
	default:
		panic("[shaXStr] x must be in [1, 256, 512]")
	}

	h.Write(str)
	hBytes := h.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}
