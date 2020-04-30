package co

import (
	"crypto/md5"
	"encoding/hex"
)

// 字符串生成md5摘要
func Md5Str(s string) string {
	data := []byte(s)
	m := md5.New()
	_, _ = m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

// 获取字符串md5值,length指定结果长度32/16
func Md5(str string, length uint8) string {
	return string(md5Str([]byte(str), length))
}

// 计算字符串的 MD5 散列值
func md5Str(str []byte, length uint8) []byte {
	var res []byte
	h := md5.New()
	h.Write(str)

	hBytes := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(dst, hBytes)
	if length > 0 && length < 32 {
		res = dst[:length]
	} else {
		res = dst
	}

	return res
}
