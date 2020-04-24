package convert

import (
	"reflect"
)

// IsEmpty 检查变量是否为空.
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// IsNil 检查变量是否空值.
func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface:
		if rv.IsNil() {
			return true
		}
	}
	return false
}

// IsBool 是否布尔值.
func IsBool(val interface{}) bool {
	return val == true || val == false
}

// IsHex 是否十六进制字符串.
func IsHex(str string) bool {
	_, err := Hex2Dec(str)
	return err == nil
}

// IsByte 变量是否字节切片.
func IsByte(val interface{}) bool {
	return Gettype(val) == "[]uint8"
}

// IsStruct 变量是否结构体.
func IsStruct(val interface{}) bool {
	r := reflectPtr(reflect.ValueOf(val))
	return r.Kind() == reflect.Struct
}

// IsInterface 变量是否接口.
func IsInterface(val interface{}) bool {
	r := reflectPtr(reflect.ValueOf(val))
	return r.Kind() == reflect.Invalid
}

// reflectPtr 获取反射的指向.
func reflectPtr(r reflect.Value) reflect.Value {
	// 如果是指针,则获取其所指向的元素
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r
}
