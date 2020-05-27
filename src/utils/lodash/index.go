package lodash

import (
	"ecode/models"
	"errors"
	"reflect"
	"strings"
)

// StructToMap struct 结构 转 map
func StructToMap(in interface{}) (models.H, error) {
	out := make(models.H)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("ToMap 只接受 structs，现在是" + v.Kind().String())
	}
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		out[fi.Name] = v.Field(i).Interface()
	}
	return out, nil
}

// IndexOf 使用反射黑科技，判断 一个 list 中是否包含指定值
func IndexOf(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

// StringsIndexOf 判断 一个 list 中是否包含指定字符串
func StringsIndexOf(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Pick 拾取 map
func Pick(in models.H, array []string) models.H {
	out := make(models.H)
	for k, v := range in {
		if StringsIndexOf(array, k) >= 0 {
			out[k] = v
		}
	}
	return out
}

// Omit 剪裁 map
func Omit(in models.H, array []string) models.H {
	out := make(models.H)
	for k, v := range in {
		if StringsIndexOf(array, k) == -1 {
			out[k] = v
		}
	}
	return out
}

// ToLowerFirstCase 首字母小写
func ToLowerFirstCase(str string) string {
	strMap := strings.Split(str, "")
	return strings.ToLower(strMap[0]) + strings.Join(strMap[1:], "")
}
