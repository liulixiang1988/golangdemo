package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id     int    `name:id;isnull:false`
	Name   string `name:name; isnull:true`
	Gender string `name:-`
}

//struct转化为对应的Value
func Struct2val(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr { //v.Kind()判断v的类型
		v = v.Elem() //返回指针指向的类型
	}

	if v.Kind() != reflect.Struct {
		panic("not struct") //如果v的类型不是struct，panic
	}
	return v
}

//获取struct的信息
func StructInfo(i interface{}) (reflect.Value, []reflect.StructField) {
	v := Struct2val(i)
	t := v.Type() //等价于reflect.TypeOf(i)

	f := make([]reflect.StructField, 0)

	for index := 0; index < t.NumField(); index++ { //如果t.Kind()是struct，则t.NumField()返回对应struct字段的数量
		field := t.Field(index) //返回struct对应字段的StructField

		//不能导出小写字段
		if field.PkgPath != "" {
			continue
		}

		//如果tag中的name为-，不导出
		//StuctTag.Get() 如果不存在对应的tag，则返回""
		if tag := field.Tag.Get("name"); tag == "-" {
			continue
		}
		f = append(f, field)
	}

	return v, f
}

//判断是否是struct
func IsStruct(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

//返回类型的名，如果是匿名结构，返回空字符串
func Name(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("not struct")
	}
	return t.Name()
}

//返回所有的字段
func Fields(i interface{}) []string {
	v, fields := StructInfo(i)

}

//判断struct中是否存在有对应的字段
func Has(i interface{}, fieldName string) bool {
	value, fields := StructInfo(i)
	for i, field := range fields {

		//如果struct中的字段还是struct，则进行迭代
		//Value.Interface()的作用是获得Value对应的值
		if IsStruct(value.Interface()) {
			if ok = Has(value.Interface(), fieldName); ok {
				return true
			}
		}
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Hello, 世界")
}
