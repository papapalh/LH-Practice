package main

import (
	"fmt"
	"reflect"
)

func main(){
	type S struct {
		F string `species:"gopher" color:"blue"`
		X string `species:"X" color:"X"`
	}
	s := S{}


	st := reflect.TypeOf(s)
	fmt.Print("TypeOf(i) go的反射接口"); fmt.Println(st)

	fmt.Print("Field(i) 返回结构体的第i个字段"); field := st.Field(1); fmt.Println(field)
	fmt.Print("NumField() 返回结构体的字段个数"); fmt.Println(st.NumField())
	fmt.Print("FieldByName() 返回该类型名为name的字段"); fmt.Println(st.FieldByName("X"))
	fmt.Print("Get() Get方法返回标签字符串中键key对应的值"); fmt.Println(field.Tag.Get("color"))
}