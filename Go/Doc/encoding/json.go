package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID     int      `json:"-"`                 //忽略该JSON字段(encode/decode都不出现这个字段)
	Name   string   `json:"myName"`            //正常JSON定义(字段名为myName)
	Name2  string   `json:"myName2,omitempty"` //字段为空，则省略(如果字段未定义，则encode/decode都不出现这个字段)
	Colors []string `json:",omitempty"`        //JSON为默认值，但是如果字段为空，则省略(如果字段未定义，则encode/decode都不出现这个字段))
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	// json_encode方法(不可以是通道/函数/复数)
	a, _ := json.Marshal(group)
	fmt.Println(string(a))

	// json_decode 解析json
	group2 := ColorGroup{}
	json.Unmarshal(a, &group2)
	fmt.Println(group2)
}
