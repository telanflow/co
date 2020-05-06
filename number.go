package co

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
)

// 将整数转换为字符串
func Int2Str(val interface{}) string {
	switch val.(type) {
	// Integers
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}
}

// 将浮点数转换为字符串,decimal为小数位数
func Float2Str(val interface{}, decimal int) string {
	switch val.(type) {
	// Floats
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', decimal, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', decimal, 64)
	// Type is not floats, return empty string
	default:
		return ""
	}
}

// 将布尔值转换为字符串
func Bool2Str(val bool) string {
	if val {
		return "true"
	}
	return "false"
}

// 将布尔值转换为整型
func Bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

// 将十进制转换为二进制
func Dec2Bin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// 将二进制转换为十进制
func Bin2Dec(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// 将十六进制字符串转换为二进制字符串
func Hex2Bin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// 将二进制字符串转换为十六进制字符串
func Bin2Hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// 将十进制转换为十六进制
func Dec2Hex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// 将十六进制转换为十进制
func Hex2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽
	return strconv.ParseInt(str[start:], 16, 0)
}

// 将十进制转换为八进制
func Dec2Oct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// 将八进制转换为十进制
func Oct2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 1 && str[0:1] == "0" {
		start = 1
	}

	return strconv.ParseInt(str[start:], 8, 0)
}

// 进制转换,在任意进制之间转换数字
// number为输入数值, frombase为原进制, tobase为结果进制
func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// 将 IPV4 的字符串互联网协议转换成长整型数字
func Ip2Long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// 将长整型转化为字符串形式带点的互联网标准格式地址(IPV4)
func Long2Ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

// 获取变量类型
func Gettype(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 强制将变量转换为字符串
func ToStr(val interface{}) string {
	//先处理其他类型
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}

	//再处理字节切片
	switch val.(type) {
	case []uint8:
		return string(val.([]uint8))
	}

	return fmt.Sprintf("%v", val)
}

// 强制将变量转换为布尔值
func ToBool(val interface{}) bool {
	switch val.(type) {
	case int:
		return (val.(int) > 0)
	case int8:
		return (val.(int8) > 0)
	case int16:
		return (val.(int16) > 0)
	case int32:
		return (val.(int32) > 0)
	case int64:
		return (val.(int64) > 0)
	case uint:
		return (val.(uint) > 0)
	case uint8:
		return (val.(uint8) > 0)
	case uint16:
		return (val.(uint16) > 0)
	case uint32:
		return (val.(uint32) > 0)
	case uint64:
		return (val.(uint64) > 0)
	case float32:
		return (val.(float32) > 0)
	case float64:
		return (val.(float64) > 0)
	case []uint8:
		return Str2Bool(string(val.([]uint8)))
	case string:
		return Str2Bool(val.(string))
	case bool:
		return val.(bool)
	default:
		return false
	}
}

// 强制将变量转换为整型;其中true或"true"为1
func ToInt(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case uint:
		return int(val.(uint))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float32:
		return int(val.(float32))
	case float64:
		return int(val.(float64))
	case []uint8:
		return Str2Int(string(val.([]uint8)))
	case string:
		return Str2Int(val.(string))
	case bool:
		return Bool2Int(val.(bool))
	default:
		return 0
	}
}

// 强制将变量转换为浮点型;其中true或"true"为1.0
func ToFloat(val interface{}) (res float64) {
	switch val.(type) {
	case int:
		res = float64(val.(int))
	case int8:
		res = float64(val.(int8))
	case int16:
		res = float64(val.(int16))
	case int32:
		res = float64(val.(int32))
	case int64:
		res = float64(val.(int64))
	case uint:
		res = float64(val.(uint))
	case uint8:
		res = float64(val.(uint8))
	case uint16:
		res = float64(val.(uint16))
	case uint32:
		res = float64(val.(uint32))
	case uint64:
		res = float64(val.(uint64))
	case float32:
		res = float64(val.(float32))
	case float64:
		res = val.(float64)
	case []uint8:
		res = Str2Float64(string(val.([]uint8)))
	case string:
		res = Str2Float64(val.(string))
	case bool:
		if val.(bool) {
			res = 1.0
		}
	}

	return
}

// 64位浮点数转字节切片
func Float64ToByte(val float64) []byte {
	bits := math.Float64bits(val)
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, bits)

	return res
}

// 64位整型转字节切片
func Int64ToByte(val int64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, uint64(val))
	return res
}

// 16进制字符串转字节切片
func Hex2Byte(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

// 16进制切片转byte切片
func HexSlice2Byte(val []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(val)))
	_, err := hex.Decode(dst, val)
	if err != nil {
		return nil
	}
	return dst
}
