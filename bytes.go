package co

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"unsafe"
)

// 将字节切片转换为字符串
// 零拷贝,不安全.效率是string([]byte{})的百倍以上,且转换量越大效率优势越明显
func Bytes2Str(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

// 字节切片转64位浮点数
func Byte2Float64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// 字节切片转64位整型
func Byte2Int64(val []byte) int64 {
	return int64(binary.BigEndian.Uint64(val))
}

// 字节切片转16进制字符串
func Byte2Hex(val []byte) string {
	return hex.EncodeToString(val)
}

// 字节切片转16进制切片
func Byte2HexSlice(val []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(val)))
	hex.Encode(dst, val)
	return dst
}
