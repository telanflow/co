package co

import (
	"reflect"
	"strings"
)

// 合并给定的2个map
// ignoreCase 过滤重复的键
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}

		dst[k] = v
	}

	return dst
}

// 将给定的map键名转换为小写
func KeyToLower(src map[string]string) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		k = strings.ToLower(k)
		newMp[k] = v
	}

	return newMp
}

// 给定key获取value, 支持递归获取
// eg "top" "top.sub"
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool) {
	if val, ok := mp[key]; ok {
		return val, true
	}

	// has sub key? eg. "top.sub"
	if !strings.ContainsRune(key, '.') {
		return nil, false
	}

	keys := strings.Split(key, ".")
	topK := keys[0]

	// find top item data based on top key
	var item interface{}
	if item, ok = mp[topK]; !ok {
		return
	}

	for _, k := range keys[1:] {
		switch tData := item.(type) {
		case map[string]string: // is simple map
			item, ok = tData[k]
			if !ok {
				return
			}
		case map[string]interface{}: // is map(decode from toml/json)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[interface{}]interface{}: // is map(decode from yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		default: // error
			ok = false
			return
		}
	}

	return item, true
}

// 获取给定map的所有键
func Keys(mp interface{}) (keys []string) {
	rftVal := reflect.ValueOf(mp)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}

	if rftVal.Kind() != reflect.Map {
		return
	}

	for _, key := range rftVal.MapKeys() {
		keys = append(keys, key.String())
	}
	return
}

// 获取给定map的所有值
func Values(mp interface{}) (values []interface{}) {
	rftTyp := reflect.TypeOf(mp)
	if rftTyp.Kind() == reflect.Ptr {
		rftTyp = rftTyp.Elem()
	}

	if rftTyp.Kind() != reflect.Map {
		return
	}

	rftVal := reflect.ValueOf(mp)
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}
