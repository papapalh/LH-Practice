package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ColorGroup struct {
	ID     int      `json:"-"`                 //忽略该JSON字段
	Name   string   `json:"myName"`            //正常JSON定义
	Name2  string   `json:"myName2,omitempty"` //字段为空，则省略
	Colors []string `json:",omitempty"`        //JSON为默认值，但是如果字段为空，则省略
}

func main(){
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	a, _ := json.Marshal(group)
	fmt.Print("Marshal(v) json_encode方法(不可以是通道/函数/复数)"); fmt.Println(string(a))

	group2 := ColorGroup{}
	b := json.NewDecoder(strings.NewReader(string(a)))
	b.Decode(&group2)
	fmt.Print("NewDecoder(i) 获取json / Decode(*v) 解析json"); fmt.Println(group2)
}