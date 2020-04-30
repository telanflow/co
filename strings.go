package co

import (
	"reflect"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

// 将字符串转换为字节切片.
// 该方法零拷贝,但不安全.它直接转换底层指针,两者指向的相同的内存,改一个另外一个也会变.
// 仅当临时需将长字符串转换且不长时间保存时可以使用.
// 转换之后若没做其他操作直接改变里面的字符,则程序会崩溃.
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic.
func Str2Bytes(s string) []byte {
	pSliceHeader := &reflect.SliceHeader{}
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pSliceHeader.Data = strHeader.Data
	pSliceHeader.Len = strHeader.Len
	pSliceHeader.Cap = strHeader.Len
	return *(*[]byte)(unsafe.Pointer(pSliceHeader))
}

// 严格将字符串转换为有符号整型.
// bitSize为类型位数,strict为是否严格检查.
func Str2IntStrict(val string, bitSize int, strict bool) int64 {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// 将字符串转换为int.其中"true", "TRUE", "True"为1.
func Str2Int(val string) (res int) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1
		return
	}

	res, _ = strconv.Atoi(val)
	return
}

// 将字符串转换为int8.
func Str2Int8(val string) int8 {
	return int8(Str2IntStrict(val, 8, false))
}

// 将字符串转换为int16.
func Str2Int16(val string) int16 {
	return int16(Str2IntStrict(val, 16, false))
}

// 将字符串转换为int32.
func Str2Int32(val string) int32 {
	return int32(Str2IntStrict(val, 32, false))
}

// 将字符串转换为int64.
func Str2Int64(val string) int64 {
	return Str2IntStrict(val, 64, false)
}

// 严格将字符串转换为无符号整型,bitSize为类型位数,strict为是否严格检查
func Str2UintStrict(val string, bitSize int, strict bool) uint64 {
	res, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// 将字符串转换为uint.
func Str2Uint(val string) uint {
	return uint(Str2UintStrict(val, 0, false))
}

// 将字符串转换为uint8.
func Str2Uint8(val string) uint8 {
	return uint8(Str2UintStrict(val, 8, false))
}

// 将字符串转换为uint16.
func Str2Uint16(val string) uint16 {
	return uint16(Str2UintStrict(val, 16, false))
}

// 将字符串转换为uint32.
func Str2Uint32(val string) uint32 {
	return uint32(Str2UintStrict(val, 32, false))
}

// 将字符串转换为uint64.
func Str2Uint64(val string) uint64 {
	return Str2UintStrict(val, 64, false)
}

// 严格将字符串转换为浮点型.
// bitSize为类型位数,strict为是否严格检查.
func Str2FloatStrict(val string, bitSize int, strict bool) float64 {
	res, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// 将字符串转换为float32.
func Str2Float32(val string) float32 {
	return float32(Str2FloatStrict(val, 32, false))
}

// 将字符串转换为float64.其中"true", "TRUE", "True"为1.0 .
func Str2Float64(val string) (res float64) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1.0
	} else {
		res = Str2FloatStrict(val, 64, false)
	}
	return
}

// 将字符串转换为布尔值.
// 1, t, T, TRUE, true, True 等字符串为真.
// 0, f, F, FALSE, false, False 等字符串为假.
func Str2Bool(val string) (res bool) {
	if val != "" {
		res, _ = strconv.ParseBool(val)
	}
	return
}

// 将[]rune转为[]byte.
func Runes2Bytes(rs []rune) []byte {
	size := 0
	for _, r := range rs {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)
	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(bs[count:], r)
	}
	return bs
}