package co

import (
	"math/rand"
	"time"
	"unicode"
)

var (
	gBytes      = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	gBytesTotal = len(gBytes)
)

// 随机获取 min - max 之间的数值
func RandInt(min, max int) int {
	if min == max {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// 随机获取 min - max 之间的数值切片
func RandIntSlice(length, min, max int) []int {
	list := make([]int, 0, length)
	for {
	LOOP:
		if len(list) == length {
			break
		}
		n := RandInt(min, max)
		for _, v := range list {
			if v == n {
				goto LOOP
			}
		}
		list = append(list, n)
	}
	return list
}

// 随机获取 min - max 之间的数值切片
func RandInt64Slice(length int, min, max int64) []int64 {
	list := make([]int64, 0, length)
	for {
	LOOP:
		if len(list) == length {
			break
		}
		n := RandInt64(min, max)
		for _, v := range list {
			if v == n {
				goto LOOP
			}
		}
		list = append(list, n)
	}
	return list
}

// 随机获取 min - max 之间的数值
func RandInt64(min, max int64) int64 {
	if min == max {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// 生成随机字符串
func RandStr(length int) string {
	return string(RandBytes(length))
}

// 生成随机字节切片
func RandBytes(length int) []byte {
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		b[i] = gBytes[rand.Intn(gBytesTotal)]
	}
	return b
}

// 生成随机字符串
func RandStrForBytes(length int, payload []byte) string {
	var (
		total = len(payload)
		b     = make([]byte, length)
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		b[i] = payload[rand.Intn(total)]
	}
	return string(b)
}

// 字符串随机转大写
func RandStrToUpper(str string, upperSize int) string {
	var (
		runes = []rune(str)
		total = len(runes)
	)
	if upperSize > total {
		upperSize = total
	}

	rand.Seed(time.Now().UnixNano())
	indexList := RandIntSlice(upperSize, 0, total-1)
	for i := 0; i < upperSize; i++ {
		index := indexList[i]
		runes[index] = unicode.ToUpper(runes[index])
	}
	return string(runes)
}

// 字符串随机转小写
func RandStrToLower(str string, lowerSize int) string {
	var (
		runes = []rune(str)
		total = len(runes)
	)
	if lowerSize > total {
		lowerSize = total
	}

	rand.Seed(time.Now().UnixNano())
	indexList := RandIntSlice(lowerSize, 0, total-1)
	for i := 0; i < lowerSize; i++ {
		index := indexList[i]
		runes[index] = unicode.ToLower(runes[index])
	}
	return string(runes)
}

// 指定字符索引转大写
func RandStrToUpperWithIndex(str string, indexList ...int) string {
	var (
		runes = []rune(str)
		total = len(indexList)
	)
	for i := 0; i < total; i++ {
		index := indexList[i]
		runes[index] = unicode.ToUpper(runes[index])
	}
	return string(runes)
}

// 指定字符索引转小写
func RandStrToLowerWithIndex(str string, indexList ...int) string {
	var (
		runes = []rune(str)
		total = len(indexList)
	)
	for i := 0; i < total; i++ {
		index := indexList[i]
		runes[index] = unicode.ToLower(runes[index])
	}
	return string(runes)
}
