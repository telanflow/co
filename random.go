package co

import (
	"math/rand"
	"time"
)

const payloadRandStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机获取 min - max 之间的数值
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

// 随机获取 min - max 之间的数值
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// 生成随机字符串
func RandStr(length int) string {
	var (
		bytes  = []byte(payloadRandStr)
		result = make([]byte, 0)
		total  = len(bytes)
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(total)])
	}
	return string(result)
}
