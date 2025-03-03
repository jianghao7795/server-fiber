package utils

import (
	"fmt"
	"reflect"
	"strings"
)

//@author: wuhao
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj any
//@return: map[string]any

func StructToMap(obj any) map[string]any {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]any)
	for i := range obj1.NumField() {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: wuhao
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []any
//@return: string

func ArrayToString(array []any) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
